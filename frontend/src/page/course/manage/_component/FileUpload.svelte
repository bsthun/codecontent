<script lang="ts">
	import { createEventDispatcher } from 'svelte'
	import { UploadIcon, CheckCircleIcon, AlertCircleIcon, ImageIcon } from '@lucide/svelte'
	import { Button } from '$/lib/shadcn/components/button'

	export type Props = {
		courseId: number
		class?: string
		onUploaded?: () => void
	}

	let { courseId, class: className, onUploaded }: Props = $props()

	const dispatch = createEventDispatcher()

	let dragOver = $state(false)
	let fileInput: HTMLInputElement
	let uploadProgress = $state(0)
	let uploadStatus = $state<'idle' | 'uploading' | 'success' | 'error'>('idle')
	let errorMessage = $state('')
	let uploading = $state(false)

	const StatusIcon = $derived(
		uploadStatus === 'success' ? CheckCircleIcon : uploadStatus === 'error' ? AlertCircleIcon : ImageIcon
	)
	const statusText = $derived(
		uploadStatus === 'uploading'
			? `Uploading... ${uploadProgress}%`
			: uploadStatus === 'success'
			? 'Photo uploaded successfully!'
			: uploadStatus === 'error'
			? errorMessage || 'Upload failed'
			: 'Drop photo to upload or click to browse'
	)
	const color = $derived(
		uploadStatus === 'uploading'
			? 'text-blue-600'
			: uploadStatus === 'success'
			? 'text-green-600'
			: uploadStatus === 'error'
			? 'text-red-600'
			: 'text-muted-foreground'
	)
	const borderClasses = $derived(
		dragOver
			? 'border-primary bg-primary/5'
			: uploadStatus === 'error'
			? 'border-destructive bg-destructive/5'
			: uploadStatus === 'success'
			? 'border-green-500 bg-green-50 dark:bg-green-950/20'
			: 'border-dashed border-muted-foreground/25 hover:border-muted-foreground/50'
	)

	const handleDragOver = (event: DragEvent) => {
		event.preventDefault()
		dragOver = true
	}

	const handleDragLeave = () => {
		dragOver = false
	}

	const handleDrop = (event: DragEvent) => {
		event.preventDefault()
		dragOver = false

		if (event.dataTransfer?.files) {
			handleFiles(event.dataTransfer.files)
		}
	}

	const handleFileInputChange = (event: Event) => {
		const input = event.target as HTMLInputElement
		if (input.files) {
			handleFiles(input.files)
		}
	}

	const handleFiles = (files: FileList) => {
		if (files.length > 1) {
			errorMessage = 'Please upload only one file at a time'
			uploadStatus = 'error'
			setTimeout(() => {
				uploadStatus = 'idle'
				errorMessage = ''
			}, 3000)
			return
		}

		const file = files[0]
		if (!file.type.startsWith('image/')) {
			errorMessage = 'Please upload an image file'
			uploadStatus = 'error'
			setTimeout(() => {
				uploadStatus = 'idle'
				errorMessage = ''
			}, 3000)
			return
		}

		uploadFile(file)
	}

	const uploadFile = (file: File) => {
		uploading = true
		uploadStatus = 'uploading'
		uploadProgress = 0

		const formData = new FormData()
		formData.append('courseId', courseId.toString())
		formData.append('file', file)

		const xhr = new XMLHttpRequest()

		xhr.upload.addEventListener('progress', (event) => {
			if (event.lengthComputable) {
				uploadProgress = Math.round((event.loaded * 100) / event.total)
			}
		})

		xhr.addEventListener('load', () => {
			if (xhr.status === 200) {
				uploadStatus = 'success'
				dispatch('uploaded')
				onUploaded?.()

				setTimeout(() => {
					uploading = false
					uploadStatus = 'idle'
					uploadProgress = 0
					if (fileInput) {
						fileInput.value = ''
					}
				}, 2000)
			} else {
				handleUploadError()
			}
		})

		xhr.addEventListener('error', () => {
			handleUploadError()
		})

		xhr.open('POST', '/api/courses/photo/upload')
		xhr.send(formData)
	}

	const handleUploadError = () => {
		uploadStatus = 'error'
		errorMessage = 'Upload failed. Please try again.'

		setTimeout(() => {
			uploading = false
			uploadStatus = 'idle'
			uploadProgress = 0
			errorMessage = ''
			if (fileInput) {
				fileInput.value = ''
			}
		}, 3000)
	}

	const triggerFileInput = () => {
		if (!uploading) {
			fileInput.click()
		}
	}
</script>

<div class={className}>
	<div
		class="group relative flex cursor-pointer items-center justify-center rounded-lg border-2 p-6 transition-all duration-200 {borderClasses}"
		role="button"
		tabindex="0"
		ondragover={handleDragOver}
		ondragleave={handleDragLeave}
		ondrop={handleDrop}
		onclick={triggerFileInput}
		onkeydown={(e) => e.key === 'Enter' && triggerFileInput()}
	>
		<input
			bind:this={fileInput}
			type="file"
			accept="image/*"
			class="hidden"
			onchange={handleFileInputChange}
		/>

		<div class="flex flex-col items-center gap-3 text-center">
			<div
				class="flex h-12 w-12 items-center justify-center rounded-full bg-muted transition-colors duration-200 group-hover:bg-primary/10"
			>
				<StatusIcon class="h-6 w-6 {color}" />
			</div>

			<div class="space-y-1">
				<p class="text-sm font-medium {color}">{statusText}</p>
				{#if uploadStatus === 'idle'}
					<p class="text-muted-foreground text-xs">PNG, JPG, GIF up to 10MB</p>
				{/if}
			</div>

			{#if uploadStatus === 'uploading'}
				<div class="w-48 space-y-2">
					<div class="bg-muted h-2 w-full overflow-hidden rounded-full">
						<div
							class="bg-primary h-full transition-all duration-300 ease-out"
							style="width: {uploadProgress}%"
						></div>
					</div>
					<p class="text-muted-foreground text-xs">{uploadProgress}% complete</p>
				</div>
			{:else if uploadStatus === 'idle'}
				<Button variant="outline" size="sm" class="gap-2">
					<UploadIcon class="h-4 w-4" />
					Choose Photo
				</Button>
			{/if}
		</div>

		{#if uploadStatus === 'success'}
			<div class="absolute -top-2 -right-2">
				<div class="bg-green-500 flex h-6 w-6 items-center justify-center rounded-full text-white">
					<CheckCircleIcon class="h-4 w-4" />
				</div>
			</div>
		{/if}
	</div>
</div>