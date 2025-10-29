# Frontend

## Guideline

- If any variable should apply in the name, it will have `#variableName#` in the name, e.g., `#entityName#IdRequest`
- Strictly follow CLAUDE.md regardless of code implemented before

## General

- Do not add comment on anything during implementation unless complex logic is involved, and the comment must be
  lowercase short text (not more than 5 words)
- Run `bun run check` after implementation
- Use `tree` to explore folder structure if needed

## Design

- Use `@lucide/svelte` for icons, import with `Icon` suffix, e.g., `import { XIcon } from '@lucide/svelte'`
- Use shadcn components from `$/lib/shadcn/components/`, e.g., `Card`, `Button`, etc.
- Always fetch https://www.shadcn-svelte.com/docs/components/button-group.md for example before implementation,
  replacing `button-group` with the component name you want to use.
- Available components: accordion, alert, alert-dialog, aspect-ratio, avatar, badge, breadcrumb, button, button-group,
  card, carousel, checkbox, collapsible, command, data-table, dialog, drawer, dropdown-menu, empty, field, input,
  input-group, item, label, menubar, pagination, popover, radio-group, resizable, scroll-area, select, separator, sheet,
  sidebar, skeleton, spinner, switch, table, tabs, textarea

## Interaction

- Use `svelte-navigator` for routing, e.g., `<Link to="..."></Link>` or
  `import { useNavigate, useParams } from 'svelte-navigator'`
- Use loading from `import Loading from '$/component/interact/Loading.svelte'` which have props container = false
- (weather to add wrapper to use in page or not), class = '', size = 'md'

## Setup

- Use setup context which have `{ profile: { id, name, email, pictureUrl, isAdmin } }` globally for user id

    ```
    import { getContext } from 'svelte'
    import type { Writable } from 'svelte/store'
    import type { Setup } from '$/util/type/setup'

    const setup = getContext<Writable<Setup>>('setup')
    ```

## Backend

- API specification is in `src/util/backend/backend.md`, follow the specification strictly
- Use `backend` utils with .then() / .catch() style like this:

    ```ts
    import { backend, catcher } from '$/util/backend'

    const fetchEnrolledCourses = () => {
    	loadingEnrolled = true
    	backend.courses
    		.courseListEnroll({
    			userId,
    			name: searchQuery,
    			limit,
    			offset,
    		})
    		.then((res) => {
    			enrolledCourses = res.data.items
    			loadingEnrolled = false
    		})
    		.catch((err) => {
    			catcher(err)
    			loadingEnrolled = false
    		})
    }
    ```

- `fetch#backendName#` function must called from `onMount` lifecycle hook

## Component

- When implementing a feature, always break down into smaller components under `_component/` folder first
- Use `$props` and `$bindable` for props implementation,

    ```
    import type { Snippet } from 'svelte'

    export type Props = {
      class?: string,
        children: Snippet,
        open: $bindable(false),
    }

    let { class: className, children, open } = $props
    ```

- Use {@render children()} for children rendering
