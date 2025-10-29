<script lang="ts">
	import { backend, catcher } from '$/util/backend'
	import { XIcon } from '@lucide/svelte'

	interface Props {
		open?: boolean
		onCourseCreated?: () => void
	}

	let { open = $bindable(false), onCourseCreated }: Props = $props()

	let formData = $state({
		name: '',
		description: ''
	})

	let isSubmitting = $state(false)

	const handleClose = () => {
		if (!isSubmitting) {
			open = false
			// Reset form
			formData = {
				name: '',
				description: ''
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
				description: formData.description.trim()
			})

			// Reset form and close dialog
			formData = {
				name: '',
				description: ''
			}
			open = false

			// Trigger refresh callback
			onCourseCreated?.()
		} catch (error) {
			catcher(error as any)
		} finally {
			isSubmitting = false
		}
	}

	const handleKeydown = (event: KeyboardEvent) => {
		if (event.key === 'Escape') {
			handleClose()
		}
	}
</script>

<svelte:window on:keydown={handleKeydown} />

{#if open}
	<div class="fixed inset-0 z-50 flex items-center justify-center">
		<!-- Backdrop -->
		<div
			class="absolute inset-0 bg-black/50 backdrop-blur-sm"
			onclick={handleClose}
			role="button"
			tabindex="0"
			onkeydown={(e) => e.key === 'Enter' && handleClose()}
		></div>

		<!-- Dialog -->
		<div class="relative bg-base-100 rounded-lg shadow-xl w-full max-w-md mx-4">
			<!-- Header -->
			<div class="flex items-center justify-between p-6 border-b border-base-300">
				<h2 class="text-xl font-semibold text-base-content">Create New Course</h2>
				<button
					class="btn btn-ghost btn-sm btn-circle"
					onclick={handleClose}
					disabled={isSubmitting}
					aria-label="Close dialog"
				>
					<XIcon class="w-4 h-4" />
				</button>
			</div>

			<!-- Form -->
			<div class="p-6">
				<form onsubmit={handleSubmit}>
					<div class="form-control mb-4">
						<label class="label" for="course-name">
							<span class="label-text font-medium">Course Name</span>
							<span class="label-text-alt text-error">*</span>
						</label>
						<input
							id="course-name"
							type="text"
							placeholder="Enter course name"
							class="input input-bordered w-full"
							bind:value={formData.name}
							required
							disabled={isSubmitting}
							maxlength="100"
						/>
						<div class="label">
							<span class="label-text-alt text-base-content/60">
								{formData.name.length}/100 characters
							</span>
						</div>
					</div>

					<div class="form-control mb-6">
						<label class="label" for="course-description">
							<span class="label-text font-medium">Description</span>
							<span class="label-text-alt text-error">*</span>
						</label>
						<textarea
							id="course-description"
							placeholder="Enter course description"
							class="textarea textarea-bordered w-full h-24 resize-none"
							bind:value={formData.description}
							required
							disabled={isSubmitting}
							maxlength="500"
						></textarea>
						<div class="label">
							<span class="label-text-alt text-base-content/60">
								{formData.description.length}/500 characters
							</span>
						</div>
					</div>

					<!-- Actions -->
					<div class="flex gap-3">
						<button
							type="button"
							class="btn btn-ghost flex-1"
							onclick={handleClose}
							disabled={isSubmitting}
						>
							Cancel
						</button>
						<button
							type="submit"
							class="btn btn-primary flex-1"
							disabled={isSubmitting || !formData.name.trim() || !formData.description.trim()}
						>
							{#if isSubmitting}
								<span class="loading loading-spinner loading-sm"></span>
							{/if}
							Create Course
						</button>
					</div>
				</form>
			</div>
		</div>
	</div>
{/if}