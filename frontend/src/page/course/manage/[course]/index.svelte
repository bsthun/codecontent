<script lang="ts">
	import { onMount } from 'svelte'
	import { ArrowLeftIcon, InfoIcon, Loader2Icon } from '@lucide/svelte'
	import { Button } from '$/lib/shadcn/components/button'
	import Container from '$/component/layout/Container.svelte'
	import { backend, catcher } from '$/util/backend'
	import type { PayloadCourseManageDetailResponse } from '$/util/backend/backend'
	import { toast } from 'svelte-sonner'
	import CourseInfoCard from '../../_component/CourseInfoCard.svelte'
	import PromptInstructionCard from '../../_component/PromptInstructionCard.svelte'
	import CourseEditDialog from '../../_component/CourseEditDialog.svelte'
	import CourseStatisticsCard from '../../_component/CourseStatisticsCard.svelte'
	import CourseEnrollmentsCard from '../../_component/CourseEnrollmentsCard.svelte'

	export type Props = {
		course: number
	}

	const { course }: Props = $props()
	let courseDetail = $state<PayloadCourseManageDetailResponse>()
	let showEditDialog = $state(false)
	let loading = $state<Record<string, boolean>>({
		page: true,
		instruction: false,
	})

	const loadCourseDetail = () => {
		backend.courses
			.manageDetail({ courseId: course })
			.then((response) => {
				courseDetail = response.data
			})
			.catch((err) => {
				catcher(err)
			})
			.finally(() => {
				loading.page = false
			})
	}

	const handleEditCourse = () => {
		showEditDialog = true
	}

	const handleCourseUpdated = (updatedCourse: any) => {
		if (courseDetail) {
			courseDetail.course = updatedCourse
		}
		toast.success('Course updated successfully')
	}

	const handleDeleteCourse = () => {
		if (!courseDetail?.course) return

		const confirmDelete = confirm(`Are you sure you want to delete "${courseDetail.course.name}"? This action cannot be undone.`)
		if (!confirmDelete) return

		backend.courses
			.courseManageDelete({ courseId: courseDetail.course.id })
			.then(() => {
				toast.success('Course deleted successfully')
				window.location.href = '/course'
			})
			.catch((err) => {
				catcher(err)
			})
	}

	const handleSaveInstruction = (instruction: string) => {
		if (!courseDetail?.course) return

		loading.instruction = true
		backend.courses
			.courseManageEdit({
				id: courseDetail.course.id,
				name: courseDetail.course.name,
				description: courseDetail.course.description,
				promptInstruction: instruction,
				token: courseDetail.course.token,
				createdAt: courseDetail.course.createdAt,
				updatedAt: courseDetail.course.updatedAt,
			})
			.then((response) => {
				if (courseDetail) {
					courseDetail.course = response.data.course
				}
				toast.success('Prompt instructions updated successfully')
			})
			.catch((err) => {
				catcher(err)
			})
			.finally(() => {
				loading.instruction = false
			})
	}

	onMount(() => {
		loadCourseDetail()
	})
</script>

<Container class="min-h-screen py-8">
	<div class="mb-8">
		<div class="mb-6 flex items-center gap-4">
			<Button class="gap-2" href="/course" size="sm" variant="ghost">
				<ArrowLeftIcon class="h-4 w-4" />
				Back to Courses
			</Button>
		</div>
		<div class="flex flex-col">
			<h1 class="text-foreground mb-2 text-4xl font-bold">Course Management</h1>
			<p class="text-muted-foreground text-lg">Manage course information and settings</p>
		</div>
	</div>

	{#if loading.page}
		<div class="flex min-h-[400px] items-center justify-center">
			<Loader2Icon class="text-primary h-8 w-8 animate-spin" />
		</div>
	{:else if !courseDetail || !courseDetail.course}
		<div class="flex min-h-[400px] flex-col items-center justify-center">
			<InfoIcon class="mb-4 h-16 w-16 text-gray-400" />
			<h3 class="mb-2 text-lg font-semibold">Course not found</h3>
			<p class="text-muted-foreground mb-4">The course you're looking for doesn't exist</p>
			<Button href="/course">
				Back to Courses
			</Button>
		</div>
	{:else}
		<div class="space-y-6">
			<div class="grid grid-cols-1 gap-6 lg:grid-cols-2">
				<CourseInfoCard
					course={courseDetail.course}
					onEdit={handleEditCourse}
					onDelete={handleDeleteCourse}
				/>

				<PromptInstructionCard
					promptInstruction={courseDetail.course.promptInstruction}
					loading={loading.instruction}
					onSave={handleSaveInstruction}
				/>
			</div>

			<CourseStatisticsCard courseData={courseDetail} />

			<CourseEnrollmentsCard courseData={courseDetail} />
		</div>
	{/if}
</Container>

<CourseEditDialog
	bind:open={showEditDialog}
	course={courseDetail?.course}
	onCourseUpdated={handleCourseUpdated}
/>