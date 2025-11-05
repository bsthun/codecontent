<script lang="ts">
	import type { PayloadCourseExtended } from '$/util/backend/backend'
	import { BookOpenIcon, UsersIcon, ImageIcon } from 'lucide-svelte'
	import { Card, CardContent, CardTitle } from '$/lib/shadcn/components/card'
	import { Button } from '$/lib/shadcn/components/button'
	import { useNavigate } from 'svelte-navigator'

	export let course: PayloadCourseExtended
	export let variant: 'enroll' | 'manage' | 'explore' = 'explore'

	const navigate = useNavigate()

	const handleClick = () => {
		if (variant === 'manage') {
			navigate(`/course/manage/${course.id}`)
		} else if (variant === 'enroll') {
			navigate(`/content/${course.id}/document`)
		} else {
			console.log('Course clicked:', course.id)
		}
	}

	const getBadgeClasses = () => {
		switch (variant) {
			case 'manage':
				return 'bg-primary text-primary-foreground'
			case 'enroll':
				return 'bg-secondary text-secondary-foreground'
			default:
				return 'bg-muted text-muted-foreground'
		}
	}

	const getButtonVariant = () => {
		switch (variant) {
			case 'explore':
				return 'default'
			case 'enroll':
				return 'secondary'
			case 'manage':
				return 'outline'
			default:
				return 'default'
		}
	}
</script>

<Card
	class="cursor-pointer shadow-lg pt-0"
	onclick={handleClick}
	onkeypress={(e) => e.key === 'Enter' && handleClick()}
	role="button"
	tabindex={0}
>
	<div class="from-primary/20 to-secondary/20 relative h-48 rounded-t-lg bg-gradient-to-br">
		{#if course.coursePhotoCount > 0}
			<div class="absolute inset-0 flex items-center justify-center">
				<ImageIcon class="text-muted-foreground h-16 w-16" />
			</div>
		{:else}
			<div class="absolute inset-0 flex items-center justify-center ">
				<BookOpenIcon class="text-muted-foreground h-16 w-16" />
			</div>
		{/if}
		{#if variant === 'manage'}
			<span class="absolute top-4 right-4 rounded-full px-2 py-1 text-xs font-medium {getBadgeClasses()}">
				Manager
			</span>
		{:else if variant === 'enroll'}
			<span class="absolute top-4 right-4 rounded-full px-2 py-1 text-xs font-medium {getBadgeClasses()}">
				Enrolled
			</span>
		{/if}
	</div>

	<CardContent>
		<CardTitle class="line-clamp-2 text-lg font-bold">
			{course.name}
		</CardTitle>
		<p class="text-muted-foreground mb-4 line-clamp-3 flex-grow text-sm">
			{course.description || 'No description available'}
		</p>

		<div class="text-muted-foreground mb-4 flex items-center gap-4 text-sm">
			<div class="flex items-center gap-1" title="Enrolled students">
				<UsersIcon class="h-4 w-4" />
				<span>{course.enrollCount}</span>
			</div>
			{#if course.courseManagerCount > 0}
				<div class="flex items-center gap-1" title="Course managers">
					<UsersIcon class="h-4 w-4" />
					<span>{course.courseManagerCount}</span>
				</div>
			{/if}
			{#if course.coursePhotoCount > 0}
				<div class="flex items-center gap-1" title="Course photos">
					<ImageIcon class="h-4 w-4" />
					<span>{course.coursePhotoCount}</span>
				</div>
			{/if}
		</div>

		<div class="flex justify-end">
			{#if variant === 'explore'}
				<Button variant={getButtonVariant()} size="sm">Enroll Now</Button>
			{:else if variant === 'enroll'}
				<Button variant={getButtonVariant()} size="sm">Continue Learning</Button>
			{:else if variant === 'manage'}
				<Button variant={getButtonVariant()} size="sm">Manage Course</Button>
			{/if}
		</div>
	</CardContent>
</Card>
