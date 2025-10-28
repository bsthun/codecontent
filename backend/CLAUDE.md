# Backend

## Guideline

- If any variable should apply in the name, it will have `#variableName#` in the name, e.g., `#entityName#IdRequest`

## General

- Always use pointer as basis
- Use `r` as receiver name for all receiver functions, e.g., `func (r *Service) UserCreate(...) ...`
- Comment format: `// * lowercase compact action`
- Use camelCase for json tags
- The project structure use one declaration per file as basis, e.g., one endpoint / procedure per file, except for
  construction which has `type Handler` / `func Handle`, `type Server interface` /  `type Service struct` /
  `func Serve` in one file.

## Endpoint

### Declaration

- Create endpoint `/endpoint/#module#/handle_#entity#_#action#.go`, e.g., `handle_user_create.go` with function name
  `func (r *Handler) HandleUserCreate(c *fiber.Ctx) error`
- Dependencies will be injected from `/endpoint/#module#/handler.go` which have only Handler struct and constructor
  function `Handle`, the handle function must chop-down arguments one per line.
- Handler file must have only handler function, one per file, other utils function / type should be declared as
  procedure and payload types.
- Always register endpoint in `/endpoint/endpoint.go`
- Endpoint handle only parsing and data flow logic, can directly call database if only for basic CRUD and use procedure
  for complex business logic.

### Structure

- Every endpoint must starts with this snippet and exact format, change only payload name:
  ```go
  // * get user claims
  l := c.Locals("l").(*jwt.Token).Claims.(*common.LoginClaims)
  
  // * parse body
  body := new(payload.#payloadName#)
  if err := c.BodyParser(body); err != nil {
    return gut.Err(false, "invalid body", err)
  }
  
  // * validate body
  if err := gut.Validate(body); err != nil {
    return err
  }
  ```

- Return error as `return gut.Err(true, "error message", err)` or `return gut.Err(false, "error message", nil)` if no
  underlying error.

### Payload

- Always use pointer as basis, see this as example
  ```go
  package example
  type Request struct {
      Name     *string `json:"name"`
	  Object   *Object `json:"object"`
      Objects  []*Object `json:"objects"`
	  Aliases  []*string `json:"aliases"`
	  Mapper   map[string]*string `json:"mapper"`
	  Selected *bool `json:"selected"`
  }
  ```

- If request has only entity id, e.g. only `UserId`, use or create a struct for it with name `#entityName#IdRequest`, as
  example:
  ```go
  package payload
  type UserIdRequest struct {
      UserId *int64 `json:"userId"`
  }
  ```

- If request has other fields, use or create a struct with name `#endpointName#Request`, as example:
  ```go
  package payload
  type ExamSubmissionCompareRequest struct {
      UserId *uint64 `json:"userId"`
      Email  *string `json:"email"`
  }
  ```

- When return result at end of endpoint, use or create a struct with name `#endpointName#Response`, as example:
  ```go
  package payload
  type ExamSubmissionCompareResponse struct {
      CompareResult []*ExamSubmissionCompareResult `json:"compareResult"`
  }
  ```

- Return only this format:
  ```go
  return c.JSON(response.Success(c, &payload.ExamQuestionSubmitResponse{
		Submission: &payload.ExamSubmission{
			Id:                submission.Id,
			ExamQuestionId:    submission.ExamQuestionId,
		},
        Result: result,
  }))
  ```

## Procedure

- Create procedure `/procedure/#module#/proc_#entity#_#action#.go`, e.g., `proc_user_create.go` with function name
  `func (r *Procedure) UserCreate(ctx context.Context, params *payload.UserCreateParams) (*payload.User, *gut.ErrorInstance)`
- Dependencies will be injected from `/procedure/#module#/procedure.go` which have `type Proc interface` and
  `type Procedure struct` and constructor function `Proceed`, the procedure function must chop-down arguments one per
  line.
- Procedure file must have only procedure function, one per file without other utils function / type, those should be
  declared as separate procedure and payload types.

## Database

- Always view `/generate/schema.sql` as basis of database structure, do not read migration files to save tokens
- Querier declaration is in `./database/postgres/#mainTableName#.sql`
- Sqlc structs is pointer by default, regardless of nullability in database
- All Id variable generated will be `Id` not `ID`

### Querier naming convention

- Use entity name as prefix, e.g., `UserCreate`, `UserGetById`, `HostDeleteById`
- Available verbs: Create, Get, Detail, Count, List, Update, Delete
- By default, every table must have 4 queriers: `#entity#Create`, `#entity#GetById`, `#entity#UpdateById`,
  `#entity#DeleteById`
- Create querier must use practice of `INSERT INTO ... VALUES (...) RETURNING *` to return created row, args will have
  all fields included except id, created_at updated_at
- Get querier tried to select * of the entity by id
- Detail querier tried to select *, with all parent relations embedded, and child relations counted.
- Count querier use practice of `SELECT COALESCE(COUNT(*), 0)::BIGINT AS #entity#_count FROM ...` to return 0 fallback
- List querier select * by default, but if field is text, omit the field and select separately to avoid large text load,
  if any of child relations is found, add count (using subquery) of child and if any parent relation is found, embed
  parent, e.g.,
  ```sql
    -- name: ClassExamListByClass :many
    SELECT sqlc.embed(exams),
    sqlc.embed(classes),
    (SELECT COALESCE(COUNT(*), 0)::BIGINT FROM exam_questions WHERE exam_questions.exam_id = exams.id) AS exam_question_count
    FROM exams
    LEFT JOIN classes ON exams.class_id = classes.id
    WHERE exams.class_id = $1
    GROUP BY exams.id, classes.id
    ORDER BY exams.created_at DESC;
  ```
- List querier (continue), if any entity should paginate and search, use this pattern:
  ```sql
    -- name: CollectionList :many
    SELECT collections.*,
    (SELECT COALESCE(COUNT(*), 0)::BIGINT FROM collection_questions WHERE collection_questions.collection_id = collections.id) AS collection_question_count
    FROM collections
    LEFT JOIN collection_questions ON collections.id = collection_questions.collection_id
    WHERE (sqlc.narg('name')::text IS NULL OR LOWER(collections.name) LIKE LOWER('%' || sqlc.narg('name') || '%'))
    GROUP BY collections.id
    ORDER BY
      CASE WHEN sqlc.narg('sort') = 'name' THEN collections.name END ASC,
      CASE WHEN sqlc.narg('sort') = 'createdAt' THEN collections.created_at END DESC
    LIMIT $1 OFFSET $2;
  ```
- Update querier use practice of `UPDATE ... SET field = COALESCE(sqlc.narg(field), field) ... RETURNING *` to allow
  partial update and return updated row
- Use `By` to indicate query by field(s), e.g., `UserGetByEmail`, `HostListByProjectId`
- Use `And` to indicate multiple fields, e.g., `UserGetByNameAndEmail`, `HostDeleteByIdAndProjectIdAndStatus`
- It must not have any querier joined with array of child relations, e.g.,
  `SELECT users.*, sqlc.embed(roles) FROM users JOIN roles ...` create separate querier instead, `RoleListByUserId`
- Example go code to call querier:
  ```go
    // * query: user create
    createdUser, err := r.database.P().UserCreate(c.Context(), psql.CreateUserParams{
        Name:  body.Name,
        Email: body.Email,
    })
    if err != nil {
        return gut.Err(false, "failed to create user", err)
    }
  ```
- Example go code to call transaction:
  ```go
    // * begin transaction
    tx, querier := r.database.Ptx(c.Context(), nil)
    defer func() {
        if r := recover(); r != nil {
            _ = tx.Rollback()
        }
    }()
    
    // * query action create
    action, err := querier.ActionCreate(c.Context(), &psql.ActionCreateParams{
        UserId:    u.UserId,
        ProjectId: body.ProjectId,
    })
    if err != nil {
        _ = tx.Rollback()
        return gut.Err(false, "failed to create action", err)
    }
  
    ...
    
    // * commit transaction
    if err := tx.Commit(); err != nil {
        return gut.Err(false, "failed to commit transaction", err)
    }
  ```
- If a field is JSONB, it will be declared in `sqlc.yml` to map to `/type/tuple` struct, so always treats them as struct
  and do not handle json marshal / unmarshal manually, for blank value, use `[]byte("{}")` as basis.

## Implementation

- Use "github.com/bsthun/gut" for mapping arrays
    ```go
    // * map organizations to items
    organizationItems, _ := gut.Iterate(organizationRows, func(org psql.GetUserOrganizationsRow) (*payload.OrganizationItem, *gut.ErrorInstance) {
        return &payload.OrganizationItem{Id: org.Id}, nil
    })
    
    // * return
    return c.JSON(response.Success(c, &payload.OrganizationList{
        Items: organizationItems
    }))
    ```
- Anything returning `*gut.ErrorInstance` must be named `er` instead of `err`. And must be handled with
  `if er != nil { return er }` without wrapping by new `gut.Err(...)`.
- Always run `make generate` after changing anything, it includes sqlc generation, code testing, and swagger generation.