<script lang="ts">
	import type { PayloadCourseExtended } from '$/util/backend/backend'
	import CourseCard from './CourseCard.svelte'
	import { ChevronLeftIcon, ChevronRightIcon } from 'lucide-svelte'

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
	<div class="flex items-center justify-between mb-4">
		<h2 class="text-2xl font-bold text-base-content">{title}</h2>
		{#if courses.length > 0}
			<div class="flex gap-2">
				<button
					class="btn btn-circle btn-sm btn-ghost"
					on:click={() => scroll('left')}
					aria-label="Scroll left"
				>
					<ChevronLeftIcon class="w-5 h-5" />
				</button>
				<button
					class="btn btn-circle btn-sm btn-ghost"
					on:click={() => scroll('right')}
					aria-label="Scroll right"
				>
					<ChevronRightIcon class="w-5 h-5" />
				</button>
			</div>
		{/if}
	</div>

	{#if loading}
		<div class="flex gap-4 overflow-hidden">
			{#each Array(4) as _}
				<div class="skeleton h-80 w-80 shrink-0"></div>
			{/each}
		</div>
	{:else if courses.length === 0}
		<div class="card bg-base-200 shadow-lg">
			<div class="card-body text-center">
				<p class="text-base-content/60">No courses available in this section.</p>
			</div>
		</div>
	{:else}
		<div
			bind:this={scrollContainer}
			class="flex gap-6 overflow-x-auto pb-4 snap-x snap-mandatory scrollbar-hide"
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
