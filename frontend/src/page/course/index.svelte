<script lang="ts">
	import { onMount } from 'svelte'
	import { backend, catcher } from '$/util/backend'
	import type { PayloadCourseExtended } from '$/util/backend/backend'
	import Container from '$/component/layout/Container.svelte'
	import CourseSection from './_component/CourseSection.svelte'
	import SearchBox from './_component/SearchBox.svelte'
	import { PlusIcon, BookOpenIcon } from 'lucide-svelte'

	let userId: any = null
	let searchQuery: string = ''

	// Course lists
	let enrolledCourses: PayloadCourseExtended[] = []
	let managedCourses: PayloadCourseExtended[] = []
	let exploreCourses: PayloadCourseExtended[] = []

	// Loading states
	let loadingEnrolled = true
	let loadingManaged = true
	let loadingExplore = true
	let loadingUser = true

	// Pagination
	const limit = 20
	const offset = 0

	onMount(() => {
		fetchUserState()
	})

	const fetchUserState = () => {
		loadingUser = true
		backend.state
			.state()
			.then((res) => {
				userId = res.data.userId
				loadingUser = false
				fetchAllCourses()
			})
			.catch((err) => {
				catcher(err)
				loadingUser = false
			})
	}

	const fetchAllCourses = () => {
		if (!userId) return
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
		backend.courses
			.courseListExplore({
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
		// Navigate to create course page (to be implemented)
		console.log('Create new course')
	}
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
			<button class="btn btn-primary gap-2 lg:btn-lg" on:click={handleNewCourse}>
				<PlusIcon class="w-5 h-5" />
				New Course
			</button>
		</div>

		<!-- Search Box -->
		<div class="max-w-2xl">
			<SearchBox value={searchQuery} on:search={handleSearch} placeholder="Search courses..." />
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
