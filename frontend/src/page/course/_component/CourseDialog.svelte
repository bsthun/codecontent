<script lang="ts">
	import { backend, catcher } from '$/util/backend'
	import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogFooter } from '$/lib/shadcn/components/dialog'
	import { Button } from '$/lib/shadcn/components/button'
	import Loading from '$/component/interact/Loading.svelte'

	interface Props {
		open?: boolean
		onCourseCreated?: () => void
	}

	let { open = $bindable(false), onCourseCreated }: Props = $props()

	let formData = $state({
		name: '',
		description: '',
	})

	let isSubmitting = $state(false)

	const handleClose = () => {
		if (!isSubmitting) {
			open = false
			// Reset form
			formData = {
				name: '',
				description: '',
			}
		}
	}

	const handleSubmit = async () => {
		if (!formData.name.trim() || !formData.description.trim()) {
			return
		}

		isSubmitting = true

		try {
			await backend.courses.courseCreate({
				name: formData.name.trim(),
				description: formData.description.trim(),
			})

			formData = {
				name: '',
				description: '',
			}
			open = false

			onCourseCreated?.()
		} catch (error) {
			catcher(error as any)
		} finally {
			isSubmitting = false
		}
	}
</script>

<Dialog bind:open>
	<DialogContent>
		<DialogHeader>
			<DialogTitle>Create New Course</DialogTitle>
		</DialogHeader>

		<form onsubmit={handleSubmit} class="space-y-6">
			<div class="space-y-2">
				<label for="course-name" class="text-sm leading-none font-medium">
					Course Name
					<span class="text-destructive ml-1">*</span>
				</label>
				<input
					id="course-name"
					type="text"
					placeholder="Enter course name"
					class="border-input bg-background ring-offset-background placeholder:text-muted-foreground focus-visible:ring-ring flex h-10 w-full rounded-md border px-3 py-2 text-sm file:border-0 file:bg-transparent file:text-sm file:font-medium focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
					bind:value={formData.name}
					required
					disabled={isSubmitting}
					maxlength="100"
				/>
				<div class="text-muted-foreground text-xs">
					{formData.name.length}/100 characters
				</div>
			</div>

			<div class="space-y-2">
				<label for="course-description" class="text-sm leading-none font-medium">
					Description
					<span class="text-destructive ml-1">*</span>
				</label>
				<textarea
					id="course-description"
					placeholder="Enter course description"
					class="border-input bg-background ring-offset-background placeholder:text-muted-foreground focus-visible:ring-ring flex min-h-[80px] w-full rounded-md border px-3 py-2 text-sm focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
					bind:value={formData.description}
					required
					disabled={isSubmitting}
					maxlength="500"
					rows="4"
				></textarea>
				<div class="text-muted-foreground text-xs">
					{formData.description.length}/500 characters
				</div>
			</div>

			<DialogFooter>
				<Button type="button" variant="ghost" onclick={handleClose} disabled={isSubmitting}>Cancel</Button>
				<Button type="submit" disabled={isSubmitting || !formData.name.trim() || !formData.description.trim()}>
					{#if isSubmitting}
						<Loading size="sm" class="mr-2" />
					{/if}
					Create Course
				</Button>
			</DialogFooter>
		</form>
	</DialogContent>
</Dialog>
