<script lang="ts">
	import { getContext, onMount } from 'svelte'
	import type { Writable } from 'svelte/store'
	import { backend, catcher } from '$/util/backend'
	import type { PayloadCourseExtended } from '$/util/backend/backend'
	import type { Setup } from '$/util/type/setup'
	import Container from '$/component/layout/Container.svelte'
	import Loading from '$/component/interact/Loading.svelte'
	import CourseSection from './_component/CourseSection.svelte'
	import SearchBox from './_component/SearchBox.svelte'
	import CourseDialog from './_component/CourseDialog.svelte'
	import { BookOpenIcon, PlusIcon } from 'lucide-svelte'
	import { Button } from '$/lib/shadcn/components/button'

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
		showCourseDialog = true
	}

	const handleCourseCreated = () => {
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
	<div class="mb-12">
		<div class="mb-6 flex flex-col gap-6 lg:flex-row lg:items-center lg:justify-between">
			<div class="flex items-center gap-4">
				<div class="bg-primary text-primary-foreground flex h-16 w-16 items-center justify-center rounded-full">
					<BookOpenIcon class="h-8 w-8" />
				</div>
				<div>
					<h1 class="text-foreground mb-2 text-4xl font-bold">Course Content</h1>
					<p class="text-muted-foreground text-lg">A personalized coding course content platform</p>
				</div>
			</div>
			<Button class="gap-2" onclick={handleNewCourse}>
				<PlusIcon class="h-5 w-5" />
				New Course
			</Button>
		</div>

		<div class="max-w-2xl">
			<SearchBox on:search={handleSearch} placeholder="Search courses..." value={searchQuery} />
		</div>
	</div>

	{#if loadingUser}
		<div class="flex h-64 items-center justify-center">
			<Loading size="lg" />
		</div>
	{:else}
		<div class="space-y-8">
			<CourseSection
				title="Enrolled Courses"
				courses={enrolledCourses}
				variant="enrolled"
				loading={loadingEnrolled}
			/>

			<CourseSection
				title="Manage Courses"
				courses={managedCourses}
				variant="managed"
				loading={loadingManaged}
			/>

			<CourseSection
				title="Explore Courses"
				courses={exploreCourses}
				variant="explore"
				loading={loadingExplore}
			/>
		</div>
	{/if}
</Container>

<CourseDialog bind:open={showCourseDialog} onCourseCreated={handleCourseCreated} />
