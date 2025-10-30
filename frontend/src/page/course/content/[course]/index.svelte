<script lang="ts">
	import { onMount } from 'svelte'
	import { ArrowLeftIcon, FileTextIcon } from '@lucide/svelte'
	import { Button } from '$/lib/shadcn/components/button'
	import Container from '$/component/layout/Container.svelte'
	import Loading from '$/component/interact/Loading.svelte'
	import { backend, catcher } from '$/util/backend'
	import type { PayloadContentInfo } from '$/util/backend/backend'
	import { getContext } from 'svelte'
	import type { Writable } from 'svelte/store'
	import type { Setup } from '$/util/type/setup'
	import ContentCard from '../_component/ContentCard.svelte'
	import FloatingAction from '$/lib/instant/FloatingAction.svelte'
	import SearchPaginate from '$/lib/instant/SearchPaginate.svelte'
	import PromptInput from '$/lib/instant/PromptInput.svelte'
	import * as Empty from '$/lib/shadcn/components/empty'
	import { toast } from 'svelte-sonner'

	export type Props = {
		course: number
	}

	const setup = getContext<Writable<Setup>>('setup')
	const { course }: Props = $props()
	const userId = $derived($setup.profile?.id || null)

	let contents = $state<PayloadContentInfo[]>([])
	let loading = $state(true)
	let searchQuery = $state('')
	let currentPage = $state(1)
	let totalItems = $state(0)
	let totalLoading = $state(false)
	const itemsPerPage = 12

	const hasContents = $derived(contents.length > 0)
	const showEmptyState = $derived(!loading && !hasContents)

	const fetchContents = (query: string, page: number) => {
		if (totalLoading) return

		totalLoading = true
		const offset = (page - 1) * itemsPerPage

		backend.content
			.contentList({
				courseId: course,
				userId: userId as any,
				title: query,
				limit: itemsPerPage,
				offset,
				sort: 'createdAt',
				order: 'desc'
			})
			.then((response) => {
				contents = response.data.items
				totalItems = response.data.count
				currentPage = page
				searchQuery = query
			})
			.catch((err) => {
				catcher(err)
				toast.error('Failed to load contents')
			})
			.finally(() => {
				loading = false
				totalLoading = false
			})
	}

	const handleSearch = (query: string, page: number) => {
		fetchContents(query, page)
	}

	const handlePaginate = (page: number) => {
		fetchContents(searchQuery, page)
	}

	const handlePromptSubmit = (prompt: string) => {
		// Mock action - in real implementation, this would call an API
		console.log('Prompt submitted:', prompt)
	}

	onMount(() => {
		fetchContents('', 1)
	})
</script>

<Container class="min-h-screen py-8">
	<!-- Header -->
	<div class="mb-8">
		<div class="mb-6 flex items-center gap-4">
			<Button class="gap-2" href="/course/manage/{course}" size="sm" variant="ghost">
				<ArrowLeftIcon class="h-4 w-4" />
				Back to Manage
			</Button>
		</div>
		<div class="flex flex-col">
			<h1 class="text-foreground mb-2 text-4xl font-bold">Course Content</h1>
			<p class="text-muted-foreground text-lg">Browse and manage course content materials</p>
		</div>
	</div>

	{#if loading}
		<div class="flex min-h-[400px] items-center justify-center">
			<Loading container={false} />
		</div>

	<!-- Empty State -->
	{:else if showEmptyState}
		<div class="flex min-h-[400px] items-center justify-center">
			<Empty.Root>
				<Empty.Header>
					<Empty.Media variant="icon">
						<FileTextIcon class="h-12 w-12 text-muted-foreground" />
					</Empty.Media>
					<Empty.Title>{searchQuery ? 'No results found' : 'No content yet'}</Empty.Title>
					<Empty.Description>
						{searchQuery ? 'Try adjusting your search terms' : 'No content has been created for this course yet'}
					</Empty.Description>
				</Empty.Header>
				{#if searchQuery}
					<Empty.Content>
						<Button onclick={() => fetchContents('', 1)} class="gap-2">
							<ArrowLeftIcon class="h-4 w-4" />
							Clear Search
						</Button>
					</Empty.Content>
				{/if}
			</Empty.Root>
		</div>

	<!-- Content Grid -->
	{:else}
		<div class="space-y-6">
			<div class="flex items-center justify-between">
				<div class="text-sm text-muted-foreground">
					{#if searchQuery}
						Found {totalItems} result{totalItems !== 1 ? 's' : ''} for "{searchQuery}"
					{:else}
						Showing {contents.length} of {totalItems} content items
					{/if}
				</div>
			</div>

			<div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
				{#each contents as content (content.id)}
					<ContentCard {content} />
				{/each}
			</div>
		</div>
	{/if}
</Container>

<!-- Floating Action Bar -->
{#if !loading}
	<FloatingAction>
		<div class="flex flex-col gap-4">
			<SearchPaginate
				bind:query={searchQuery}
				bind:currentPage={currentPage}
				bind:totalItems={totalItems}
				{itemsPerPage}
				placeholder="Search content..."
				onSearch={handleSearch}
				onPaginate={handlePaginate}
			/>
			<PromptInput
				placeholder="Prompt new content..."
				onSubmit={handlePromptSubmit}
			/>
		</div>
	</FloatingAction>
{/if}