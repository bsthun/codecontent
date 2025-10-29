# Frontend

## Guideline

- If any variable should apply in the name, it will have `#variableName#` in the name, e.g., `#entityName#IdRequest`

## General

- Do not add comment during implementation unless necessary
- Run `bun run check` after implementation

## Design

- Use `@lucide/svelte` for icons, import with `Icon` suffix, e.g., `import { XIcon } from '@lucide/svelte'`
- Use `daisyui` for components library, see list of available components at https://daisyui.com/components/ which crawl
  https://raw.githubusercontent.com/saadeghi/daisyui/refs/heads/master/packages/docs/src/routes/(routes)/components/button/+page.md
  for full component documentation, change `button` in to other component as needed

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
        backend.courses.courseListEnroll({
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

- Use `$props` store for all bindable props, e.g., `bind:open={$props.open}`
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