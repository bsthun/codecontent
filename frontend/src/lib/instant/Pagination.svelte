<script lang="ts">
	import { ChevronLeftIcon, ChevronRightIcon } from '@lucide/svelte'
	import { Button } from '$/lib/shadcn/components/button'

	export type Props = {
		currentPage: number
		totalPages: number
		totalItems?: number
		itemsPerPage?: number
		class?: string
		onPageChange?: (page: number) => void
	}

	let {
		currentPage,
		totalPages,
		totalItems,
		itemsPerPage = 12,
		class: className,
		onPageChange
	}: Props = $props()

	const hasNextPage = $derived(currentPage < totalPages)
	const hasPrevPage = $derived(currentPage > 1)
	const startItem = $derived(totalItems ? (currentPage - 1) * itemsPerPage + 1 : null)
	const endItem = $derived(totalItems ? Math.min(currentPage * itemsPerPage, totalItems) : null)

	const handlePageChange = (page: number) => {
		if (page >= 1 && page <= totalPages && page !== currentPage) {
			onPageChange?.(page)
		}
	}

	const getVisiblePages = () => {
		const delta = 2
		const range: number[] = []
		const rangeWithDots: (number | string)[] = []
		let l: number | undefined

		for (let i = 1; i <= totalPages; i++) {
			if (i === 1 || i === totalPages || (i >= currentPage - delta && i <= currentPage + delta)) {
				range.push(i)
			}
		}

		range.forEach((i) => {
			if (l) {
				if (i - l === 2) {
					rangeWithDots.push(l + 1)
				} else if (i - l !== 1) {
					rangeWithDots.push('...')
				}
			}
			rangeWithDots.push(i)
			l = i
		})

		return rangeWithDots
	}
</script>

<div class={className}>
	{#if totalItems && itemsPerPage}
		<div class="text-muted-foreground mb-4 text-sm">
			Showing {startItem}-{endItem} of {totalItems} photos
		</div>
	{/if}

	<div class="flex items-center justify-center gap-1">
		<Button
			variant="outline"
			size="sm"
			disabled={!hasPrevPage}
			onclick={() => handlePageChange(currentPage - 1)}
			class="gap-1"
		>
			<ChevronLeftIcon class="h-4 w-4" />
			Previous
		</Button>

		<div class="flex items-center gap-1">
			{#each getVisiblePages() as page}
				{#if page === '...'}
					<span class="px-3 py-1 text-muted-foreground">...</span>
				{:else}
					<Button
						variant={page === currentPage ? 'default' : 'outline'}
						size="sm"
						onclick={() => handlePageChange(page as number)}
						class="min-w-[40px]"
					>
						{page}
					</Button>
				{/if}
			{/each}
		</div>

		<Button
			variant="outline"
			size="sm"
			disabled={!hasNextPage}
			onclick={() => handlePageChange(currentPage + 1)}
			class="gap-1"
		>
			Next
			<ChevronRightIcon class="h-4 w-4" />
		</Button>
	</div>
</div>