<script lang="ts">
	import { Card, CardHeader, CardTitle, CardDescription, CardContent } from '$/lib/shadcn/components/card'
	import type { PayloadCourseManageDetailResponse } from '$/util/backend/backend'

	export type Props = {
		courseData: PayloadCourseManageDetailResponse
		class?: string
	}

	let { class: className, courseData }: Props = $props()

	const recentEnrollments = courseData.enrollList?.slice(0, 5) || []
</script>

{#if recentEnrollments.length > 0}
	<Card class={className}>
		<CardHeader>
			<CardTitle>Recent Enrollments</CardTitle>
			<CardDescription>Students who have recently enrolled in this course</CardDescription>
		</CardHeader>
		<CardContent>
			<div class="space-y-4">
				{#each recentEnrollments as enrollment}
					<div class="flex items-center gap-4">
						<div class="bg-muted flex h-10 w-10 items-center justify-center rounded-full">
							<span class="text-sm font-medium">
								{enrollment.userFirstname ? enrollment.userFirstname[0].toUpperCase() : 'U'}
							</span>
						</div>
						<div class="flex-1">
							<p class="font-medium">
								{enrollment.userFirstname} {enrollment.userLastname}
							</p>
							<p class="text-muted-foreground text-sm">{enrollment.userEmail}</p>
						</div>
						<div class="text-right">
							<p class="text-muted-foreground text-sm">
								{new Date(enrollment.createdAt).toLocaleDateString()}
							</p>
							<p class="text-muted-foreground text-xs">
								{enrollment.contentCount || 0} items completed
							</p>
						</div>
					</div>
				{/each}
			</div>
		</CardContent>
	</Card>
{/if}