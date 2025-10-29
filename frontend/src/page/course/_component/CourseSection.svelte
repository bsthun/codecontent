<script lang="ts">
	import type { PayloadCourseExtended } from '$/util/backend/backend'
	import CourseCard from './CourseCard.svelte'
	import Loading from '$/component/interact/Loading.svelte'
	import { ChevronLeftIcon, ChevronRightIcon } from 'lucide-svelte'
	import { Button } from '$/lib/shadcn/components/button'
	import { Empty } from '$/lib/shadcn/components/empty'

	export let title: string
	export let courses: PayloadCourseExtended[]
	export let variant: 'enrolled' | 'managed' | 'explore' = 'explore'
	export let loading: boolean = false

	let scrollContainer: HTMLDivElement

	const scroll = (direction: 'left' | 'right') => {
		if (scrollContainer) {
			const scrollAmount = scrollContainer.clientWidth * 0.8
			scrollContainer.scrollBy({
				left: direction === 'left' ? -scrollAmount : scrollAmount,
				behavior: 'smooth',
			})
		}
	}
</script>

<div class="mb-12">
	<div class="mb-4 flex items-center justify-between">
		<h2 class="text-foreground text-2xl font-bold">{title}</h2>
		{#if courses.length > 0}
			<div class="flex gap-2">
				<Button
					variant="ghost"
					size="icon"
					class="h-8 w-8"
					onclick={() => scroll('left')}
					aria-label="Scroll left"
				>
					<ChevronLeftIcon class="h-4 w-4" />
				</Button>
				<Button
					variant="ghost"
					size="icon"
					class="h-8 w-8"
					onclick={() => scroll('right')}
					aria-label="Scroll right"
				>
					<ChevronRightIcon class="h-4 w-4" />
				</Button>
			</div>
		{/if}
	</div>

	{#if loading}
		<Loading container={true} />
	{:else if courses.length === 0}
		<Empty>
			<p class="text-muted-foreground">No courses available in this section.</p>
		</Empty>
	{:else}
		<div
			bind:this={scrollContainer}
			class="scrollbar-hide flex snap-x snap-mandatory gap-6 overflow-x-auto pb-4 px-2"
			style="scroll-behavior: smooth;"
		>
			{#each courses as course (course.id)}
				<div class="w-80 shrink-0 snap-start">
					<CourseCard {course} {variant} />
				</div>
			{/each}
		</div>
	{/if}
</div>

<style>
	.scrollbar-hide::-webkit-scrollbar {
		display: none;
	}

	.scrollbar-hide {
		-ms-overflow-style: none;
		scrollbar-width: none;
	}
</style>
