<script lang="ts">
	import { onMount } from 'svelte'
	import { UploadIcon, ImageIcon, ZoomInIcon, DownloadIcon, Trash2Icon } from '@lucide/svelte'
	import { Card, CardHeader, CardTitle, CardContent } from '$/lib/shadcn/components/card'
	import { Button } from '$/lib/shadcn/components/button'
	import { Dialog, DialogContent, DialogTrigger } from '$/lib/shadcn/components/dialog'
	import { backend, catcher } from '$/util/backend'
	import type { PayloadCoursePhoto } from '$/util/backend/backend'
	import Pagination from '$/lib/instant/Pagination.svelte'
	import FileUpload from './FileUpload.svelte'
	import Loading from '$/component/interact/Loading.svelte'
	import { toast } from 'svelte-sonner'

	export type Props = {
		courseId: number
		class?: string
	}

	let { courseId, class: className }: Props = $props()

	let photos = $state<PayloadCoursePhoto[]>([])
	let loading = $state(true)
	let currentPage = $state(1)
	let totalPages = $state(1)
	let totalItems = $state(0)
	const itemsPerPage = 12
	let selectedPhoto = $state<PayloadCoursePhoto | null>(null)
	let showPhotoDialog = $state(false)
	let showUploadDialog = $state(false)

	const fetchPhotos = () => {
		loading = true
		backend.courses
			.photoList({
				courseId,
				limit: itemsPerPage,
				offset: (currentPage - 1) * itemsPerPage,
				sort: 'createdAt',
				order: 'desc',
				title: ''
			})
			.then((response) => {
				photos = response.data.items
				totalItems = response.data.count
				totalPages = Math.ceil(totalItems / itemsPerPage)
			})
			.catch((err) => {
				catcher(err)
			})
			.finally(() => {
				loading = false
			})
	}

	const handlePageChange = (page: number) => {
		currentPage = page
		fetchPhotos()
	}

	const handlePhotoUploaded = () => {
		showUploadDialog = false
		currentPage = 1
		fetchPhotos()
		toast.success('Photo uploaded successfully!')
	}

	const handlePhotoClick = (photo: PayloadCoursePhoto) => {
		selectedPhoto = photo
		showPhotoDialog = true
	}

	const downloadPhoto = (photo: PayloadCoursePhoto) => {
		const link = document.createElement('a')
		link.href = photo.photoUrl
		link.download = photo.title || `photo-${photo.id}`
		document.body.appendChild(link)
		link.click()
		document.body.removeChild(link)
	}

	const deletePhoto = async (photo: PayloadCoursePhoto) => {
		if (!confirm('Are you sure you want to delete this photo?')) return

		try {
			// Note: You'll need to add the delete endpoint to backend API
			await fetch(`/api/courses/photo/${photo.id}`, {
				method: 'DELETE'
			})
			toast.success('Photo deleted successfully!')
			fetchPhotos()
		} catch (error) {
			toast.error('Failed to delete photo')
			console.error('Delete photo error:', error)
		}
	}

	onMount(() => {
		fetchPhotos()
	})

	const hasPhotos = $derived(photos.length > 0)
	const showEmptyState = $derived(!loading && !hasPhotos)
</script>

<Card class={className}>
	<CardHeader>
		<div class="flex items-center justify-between">
			<div class="flex items-center gap-2">
				<ImageIcon class="h-5 w-5 text-primary" />
				<CardTitle>Course Photos</CardTitle>
			</div>
			<Dialog bind:open={showUploadDialog}>
				<DialogTrigger>
					<Button variant="outline" size="sm" class="gap-2">
						<UploadIcon class="h-4 w-4" />
						Upload
					</Button>
				</DialogTrigger>
				<DialogContent class="max-w-md">
					<div class="space-y-4">
						<div class="text-center">
							<h3 class="text-lg font-semibold">Upload Photo</h3>
							<p class="text-muted-foreground text-sm">Add a new photo to your course gallery</p>
						</div>
						<FileUpload courseId={courseId} onUploaded={handlePhotoUploaded} />
					</div>
				</DialogContent>
			</Dialog>
		</div>
	</CardHeader>

	<CardContent>
		<div class="space-y-6">
			{#if loading}
				<div class="flex min-h-[300px] items-center justify-center">
					<Loading container={false} />
				</div>
			{:else if showEmptyState}
				<div class="flex min-h-[300px] flex-col items-center justify-center text-center">
					<div class="bg-muted/50 mb-4 flex h-16 w-16 items-center justify-center rounded-full">
						<ImageIcon class="h-8 w-8 text-muted-foreground" />
					</div>
					<h3 class="text-muted-foreground mb-2 text-lg font-medium">No photos yet</h3>
					<p class="text-muted-foreground mb-4 text-sm">Upload your first photo to get started</p>
					<Button onclick={() => (showUploadDialog = true)} class="gap-2">
						<UploadIcon class="h-4 w-4" />
						Upload First Photo
					</Button>
				</div>
			{:else}
				<div class="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4">
					{#each photos as photo (photo.id)}
						<button
							type="button"
							class="group relative aspect-square overflow-hidden rounded-lg border bg-muted/20 transition-all duration-200 hover:shadow-md hover:ring-2 hover:ring-primary/20 cursor-pointer"
							onclick={() => handlePhotoClick(photo)}
							onkeydown={(e) => e.key === 'Enter' && handlePhotoClick(photo)}
						>
							<img
								src={photo.photoUrl}
								alt={photo.title || `Course photo ${photo.id}`}
								class="h-full w-full object-cover transition-transform duration-300 group-hover:scale-105 pointer-events-none"
							/>

							<div class="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent opacity-0 transition-opacity duration-200 group-hover:opacity-100">
								<div class="absolute bottom-0 left-0 right-0 p-3">
									<div class="flex items-center justify-between">
										<div class="flex-1">
											<p class="text-white text-xs font-medium truncate">
												{photo.title || `Photo ${photo.id}`}
											</p>
											{#if photo.description}
												<p class="text-white/80 text-xs truncate">
													{photo.description}
												</p>
											{/if}
										</div>
										<div class="flex gap-1">
											<div
												role="button"
												tabindex="0"
												class="hover:bg-black/40 flex h-7 w-7 cursor-pointer items-center justify-center rounded-md bg-black/20 text-white border-0 transition-colors"
												onclick={() => {
													selectedPhoto = photo
													showPhotoDialog = true
												}}
												onkeydown={(e) => e.key === 'Enter' && (selectedPhoto = photo) && (showPhotoDialog = true)}
											>
												<ZoomInIcon class="h-3 w-3" />
											</div>
											<div
												role="button"
												tabindex="0"
												class="hover:bg-black/40 flex h-7 w-7 cursor-pointer items-center justify-center rounded-md bg-black/20 text-white border-0 transition-colors"
												onclick={() => downloadPhoto(photo)}
												onkeydown={(e) => e.key === 'Enter' && downloadPhoto(photo)}
											>
												<DownloadIcon class="h-3 w-3" />
											</div>
											<div
												role="button"
												tabindex="0"
												class="hover:bg-red-600 flex h-7 w-7 cursor-pointer items-center justify-center rounded-md bg-red-500/80 text-white border-0 transition-colors"
												onclick={() => deletePhoto(photo)}
												onkeydown={(e) => e.key === 'Enter' && deletePhoto(photo)}
											>
												<Trash2Icon class="h-3 w-3" />
											</div>
										</div>
									</div>
								</div>
							</div>
						</button>
					{/each}
				</div>

				{#if totalPages > 1}
					<Pagination
						currentPage={currentPage}
						totalPages={totalPages}
						totalItems={totalItems}
						itemsPerPage={itemsPerPage}
						onPageChange={handlePageChange}
					/>
				{/if}
			{/if}
		</div>
	</CardContent>
</Card>

<!-- Photo Preview Dialog -->
<Dialog bind:open={showPhotoDialog}>
	<DialogContent class="max-w-4xl">
		<div class="space-y-4">
			<img
				src={selectedPhoto?.photoUrl}
				alt={selectedPhoto?.title || 'Course photo'}
				class="w-full rounded-lg"
			/>
			<div class="space-y-2">
				<h3 class="text-lg font-semibold">
					{selectedPhoto?.title || `Photo ${selectedPhoto?.id}`}
				</h3>
				{#if selectedPhoto?.description}
					<p class="text-muted-foreground">
						{selectedPhoto.description}
					</p>
				{/if}
			</div>
		</div>
	</DialogContent>
</Dialog>