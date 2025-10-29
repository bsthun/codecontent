<script lang="ts">
	import { getContext, onMount } from 'svelte'
	import type { Writable } from 'svelte/store'
	import { backend, catcher } from '$/util/backend'
	import type { PayloadCourseExtended } from '$/util/backend/backend'
	import type { Setup } from '$/util/type/setup'
	import Container from '$/component/layout/Container.svelte'
	import CourseSection from './_component/CourseSection.svelte'
	import SearchBox from './_component/SearchBox.svelte'
	import CourseDialog from './_component/CourseDialog.svelte'
	import { BookOpenIcon, PlusIcon } from 'lucide-svelte'

	const setup = getContext<Writable<Setup>>('setup')
	let userId: any = null
	let searchQuery = $state('')
	let showCourseDialog = $state(false)

	// Course lists
	let enrolledCourses = $state<PayloadCourseExtended[]>([])
	let managedCourses = $state<PayloadCourseExtended[]>([])
	let exploreCourses = $state<PayloadCourseExtended[]>([])

	// Loading states
	let loadingEnrolled = $state(true)
	let loadingManaged = $state(true)
	let loadingExplore = $state(true)
	let loadingUser = $state(true)

	// Pagination
	const limit = 20
	const offset = 0

	const fetchAllCourses = () => {
		fetchEnrolledCourses()
		fetchManagedCourses()
		fetchExploreCourses()
	}

	const fetchEnrolledCourses = () => {
		loadingEnrolled = true
		backend.courses
			.courseListEnroll({
				userId,
				name: searchQuery,
				limit,
				offset,
			})
			.then((res) => {
				enrolledCourses = res.data.items
				loadingEnrolled = false
			})
			.catch((err) => {
				catcher(err)
				loadingEnrolled = false
			})
	}

	const fetchManagedCourses = () => {
		loadingManaged = true
		backend.courses
			.courseListManager({
				userId,
				name: searchQuery,
				limit,
				offset,
			})
			.then((res) => {
				managedCourses = res.data.items
				loadingManaged = false
			})
			.catch((err) => {
				catcher(err)
				loadingManaged = false
			})
	}

	const fetchExploreCourses = () => {
		loadingExplore = true
		backend.courses.courseListExplore({
			userId,
			name: searchQuery,
			limit,
			offset,
		})
			.then((res) => {
				exploreCourses = res.data.items
				loadingExplore = false
			})
			.catch((err) => {
				catcher(err)
				loadingExplore = false
			})
	}

	const handleSearch = (event: CustomEvent<string>) => {
		searchQuery = event.detail
		fetchAllCourses()
	}

	const handleNewCourse = () => {
		showCourseDialog = true
	}

	const handleCourseCreated = () => {
		// Refresh all course lists to show the new course
		fetchAllCourses()
	}

	onMount(() => {
		setup.subscribe((setupData) => {
			if (setupData?.profile?.id) {
				userId = setupData.profile.id
				loadingUser = false
				fetchAllCourses()
			}
		})
	})
</script>

<Container class="min-h-screen py-8">
	<!-- Header Section -->
	<div class="mb-12">
		<div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-6 mb-6">
			<div class="flex items-center gap-4">
				<div class="avatar placeholder">
					<div class="bg-primary text-primary-content rounded-full w-16">
						<BookOpenIcon class="w-8 h-8" />
					</div>
				</div>
				<div>
					<h1 class="text-4xl font-bold text-base-content mb-2">Course Content</h1>
					<p class="text-base-content/70 text-lg">A personalized coding course content platform</p>
				</div>
			</div>
			<button class="btn btn-primary gap-2 lg:btn-lg" onclick={handleNewCourse}>
				<PlusIcon class="w-5 h-5" />
				New Course
			</button>
		</div>

		<!-- Search Box -->
		<div class="max-w-2xl">
			<SearchBox on:search={handleSearch} placeholder="Search courses..." value={searchQuery} />
		</div>
	</div>

	{#if loadingUser}
		<div class="flex justify-center items-center h-64">
			<span class="loading loading-spinner loading-lg"></span>
		</div>
	{:else}
		<!-- Course Sections -->
		<div class="space-y-8">
			<!-- Enrolled Courses Section -->
			<CourseSection
				title="My Enrolled Courses"
				courses={enrolledCourses}
				variant="enrolled"
				loading={loadingEnrolled}
			/>

			<!-- Managed Courses Section -->
			<CourseSection
				title="Courses I Manage"
				courses={managedCourses}
				variant="managed"
				loading={loadingManaged}
			/>

			<!-- Explore Section -->
			<CourseSection
				title="Explore Courses"
				courses={exploreCourses}
				variant="explore"
				loading={loadingExplore}
			/>
		</div>
	{/if}
</Container>

<!-- Course Creation Dialog -->
<CourseDialog bind:open={showCourseDialog} onCourseCreated={handleCourseCreated} />
