<script lang="ts">
	import { Card, CardHeader, CardTitle, CardContent } from '$/lib/shadcn/components/card'
	import { Button } from '$/lib/shadcn/components/button'
	import { UsersIcon, FileTextIcon } from '@lucide/svelte'
	import type { PayloadCourseManageDetailResponse } from '$/util/backend/backend'

	export type Props = {
		courseData: PayloadCourseManageDetailResponse
		class?: string
	}

	let { class: className, courseData }: Props = $props()

	const totalSections = courseData.contentList?.reduce((sum: number, content: any) => sum + (content.contentSectionCount || 0), 0) || 0
</script>

<div class={className}>
	<div class="grid grid-cols-1 gap-6 lg:grid-cols-3">
		<Card>
			<CardHeader>
				<div class="flex items-center justify-between">
					<CardTitle class="text-lg">Enrolled Students</CardTitle>
					<Button href="/course/manage/{courseData.course.id}/enroll" variant="outline" size="sm" class="gap-1">
						<UsersIcon class="h-3 w-3" />
						View
					</Button>
				</div>
			</CardHeader>
			<CardContent>
				<div class="text-3xl font-bold">
					{courseData.enrollList?.length || 0}
				</div>
				<p class="text-muted-foreground text-sm">Total enrolled</p>
			</CardContent>
		</Card>

		<Card>
			<CardHeader>
				<div class="flex items-center justify-between">
					<CardTitle class="text-lg">Contents</CardTitle>
					<Button href="/course/content/{courseData.course.id}" variant="outline" size="sm" class="gap-1">
						<FileTextIcon class="h-3 w-3" />
						View
					</Button>
				</div>
			</CardHeader>
			<CardContent>
				<div class="text-3xl font-bold">
					{courseData.contentList?.length || 0}
				</div>
				<p class="text-muted-foreground text-sm">Total content items</p>
			</CardContent>
		</Card>

		<Card>
			<CardHeader>
				<CardTitle class="text-lg">Sections</CardTitle>
			</CardHeader>
			<CardContent>
				<div class="text-3xl font-bold">
					{totalSections}
				</div>
				<p class="text-muted-foreground text-sm">Total sections</p>
			</CardContent>
		</Card>
	</div>
</div>