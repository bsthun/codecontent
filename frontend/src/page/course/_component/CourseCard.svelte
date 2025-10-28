<script lang="ts">
	import type { PayloadCourseExtended } from '$/util/backend/backend'
	import { BookOpenIcon, UsersIcon, ImageIcon } from 'lucide-svelte'

	export let course: PayloadCourseExtended
	export let variant: 'enrolled' | 'managed' | 'explore' = 'explore'

	const getVariantClasses = () => {
		switch (variant) {
			case 'enrolled':
				return 'card-bordered hover:shadow-2xl'
			case 'managed':
				return 'card-bordered hover:shadow-2xl'
			case 'explore':
				return 'card-bordered hover:shadow-2xl'
			default:
				return 'card-bordered'
		}
	}

	const handleClick = () => {
		// Navigate to course detail page (to be implemented)
		console.log('Course clicked:', course.id)
	}
</script>

<div
	class="card bg-base-100 shadow-lg transition-all duration-300 hover:scale-105 cursor-pointer {getVariantClasses()}"
	on:click={handleClick}
	on:keypress={(e) => e.key === 'Enter' && handleClick()}
	role="button"
	tabindex="0"
>
	<figure class="relative h-48 bg-gradient-to-br from-primary/20 to-secondary/20">
		{#if course.coursePhotoCount > 0}
			<div class="absolute inset-0 flex items-center justify-center">
				<ImageIcon class="w-16 h-16 text-base-content/30" />
			</div>
		{:else}
			<div class="absolute inset-0 flex items-center justify-center">
				<BookOpenIcon class="w-16 h-16 text-base-content/30" />
			</div>
		{/if}
		{#if variant === 'managed'}
			<div class="badge badge-primary absolute top-4 right-4">Manager</div>
		{:else if variant === 'enrolled'}
			<div class="badge badge-success absolute top-4 right-4">Enrolled</div>
		{/if}
	</figure>

	<div class="card-body p-4">
		<h2 class="card-title text-lg font-bold line-clamp-2">
			{course.name}
		</h2>
		<p class="text-sm text-base-content/70 line-clamp-3 flex-grow">
			{course.description || 'No description available'}
		</p>

		<div class="flex items-center gap-4 mt-2 text-sm text-base-content/60">
			<div class="flex items-center gap-1" title="Enrolled students">
				<UsersIcon class="w-4 h-4" />
				<span>{course.enrollCount}</span>
			</div>
			{#if course.courseManagerCount > 0}
				<div class="flex items-center gap-1" title="Course managers">
					<UsersIcon class="w-4 h-4" />
					<span>{course.courseManagerCount} managers</span>
				</div>
			{/if}
			{#if course.coursePhotoCount > 0}
				<div class="flex items-center gap-1" title="Course photos">
					<ImageIcon class="w-4 h-4" />
					<span>{course.coursePhotoCount}</span>
				</div>
			{/if}
		</div>

		<div class="card-actions justify-end mt-4">
			{#if variant === 'explore'}
				<button class="btn btn-primary btn-sm">Enroll Now</button>
			{:else if variant === 'enrolled'}
				<button class="btn btn-secondary btn-sm">Continue Learning</button>
			{:else if variant === 'managed'}
				<button class="btn btn-accent btn-sm">Manage Course</button>
			{/if}
		</div>
	</div>
</div>
