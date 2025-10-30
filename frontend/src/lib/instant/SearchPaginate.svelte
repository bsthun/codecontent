<script lang="ts">
	import { createEventDispatcher } from 'svelte'
	import { Button } from '$/lib/shadcn/components/button'
	import { Input } from '$/lib/shadcn/components/input'
	import * as Pagination from '$/lib/shadcn/components/pagination'
	import { ChevronLeftIcon, ChevronRightIcon, SearchIcon } from '@lucide/svelte'

	export type Props = {
		query?: string
		currentPage?: number
		totalItems?: number
		itemsPerPage?: number
		placeholder?: string
		class?: string
		onSearch?: (query: string, page: number) => void
		onPaginate?: (page: number) => void
	}

	let {
		query = $bindable(''),
		currentPage = $bindable(1),
		totalItems = $bindable(0),
		itemsPerPage = 12,
		placeholder = 'Search...',
		class: className,
		onSearch,
		onPaginate
	}: Props = $props()

	const dispatch = createEventDispatcher<{
		search: { query: string; page: number }
		paginate: { page: number }
	}>()

	const totalPages = $derived(Math.ceil(totalItems / itemsPerPage))
	const hasNext = $derived(currentPage < totalPages)
	const hasPrevious = $derived(currentPage > 1)

	let pageInputValue = $state(currentPage.toString())

	$effect(() => {
		pageInputValue = currentPage.toString()
	})

	const handleSearch = () => {
		onSearch?.(query, 1)
		dispatch('search', { query, page: 1 })
	}

	const handleSearchKeydown = (event: KeyboardEvent) => {
		if (event.key === 'Enter') {
			handleSearch()
		}
	}

	const goToPage = (page: number) => {
		if (page >= 1 && page <= totalPages) {
			onPaginate?.(page)
			dispatch('paginate', { page })
		}
	}

	const handlePageInputKeydown = (event: KeyboardEvent) => {
		if (event.key === 'Enter') {
			const page = parseInt(pageInputValue, 10)
			if (!isNaN(page) && page >= 1 && page <= totalPages) {
				goToPage(page)
			} else {
				pageInputValue = currentPage.toString()
			}
		}
	}

	const handlePageInputBlur = () => {
		const page = parseInt(pageInputValue, 10)
		if (isNaN(page) || page < 1 || page > totalPages) {
			pageInputValue = currentPage.toString()
		}
	}

	const goToPrevious = () => {
		if (hasPrevious) {
			goToPage(currentPage - 1)
		}
	}

	const goToNext = () => {
		if (hasNext) {
			goToPage(currentPage + 1)
		}
	}

	</script>

<div class="flex items-center gap-4" class:className>
	<!-- Search Input -->
	<div class="relative">
		<SearchIcon class="text-muted-foreground absolute top-1/2 left-3 h-4 w-4 -translate-y-1/2" />
		<Input
			bind:value={query}
			class="w-64 pl-10"
			onkeydown={handleSearchKeydown}
			{placeholder}
		/>
	</div>

		<Pagination.Root count={totalItems} perPage={itemsPerPage} page={currentPage}>
			{#snippet children({ })}
				<div class="flex items-center gap-1">
					<Button disabled={!hasPrevious} onclick={goToPrevious} variant="ghost" size="sm">
						<ChevronLeftIcon class="h-4 w-4" />
					</Button>
					<Button disabled={!hasNext} onclick={goToNext} variant="ghost" size="sm">
						<ChevronRightIcon class="h-4 w-4" />
					</Button>

					<span class="text-muted-foreground">Page</span>
					<Input
						bind:value={pageInputValue}
						class="w-16 text-center"
						onblur={handlePageInputBlur}
						onkeydown={handlePageInputKeydown}
					/>
					<span class="text-muted-foreground">of {totalPages}</span>
				</div>
			{/snippet}
		</Pagination.Root>
</div>