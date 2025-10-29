<script lang="ts">
	import { Edit2Icon, Trash2Icon } from '@lucide/svelte'
	import { Card, CardHeader, CardTitle, CardDescription, CardContent } from '$/lib/shadcn/components/card'
	import { Button } from '$/lib/shadcn/components/button'
	import type { PayloadCourse } from '$/util/backend/backend'

	export type Props = {
		course: PayloadCourse
		class?: string
		onEdit?: () => void
		onDelete?: () => void
	}

	let { class: className, course, onEdit, onDelete }: Props = $props()
</script>

<Card class={className}>
	<CardHeader>
		<div class="flex items-center justify-between">
			<div>
				<CardTitle class="text-xl">{course.name}</CardTitle>
				<CardDescription class="mt-2">{course.description}</CardDescription>
			</div>
			<div class="flex gap-2">
				<Button variant="outline" size="sm" onclick={onEdit}>
					<Edit2Icon class="h-4 w-4" />
				</Button>
				<Button variant="outline" size="sm" onclick={onDelete}>
					<Trash2Icon class="h-4 w-4" />
				</Button>
			</div>
		</div>
	</CardHeader>
	<CardContent>
		<div class="space-y-4">
			<div>
				<h4 class="text-sm font-medium text-muted-foreground">Course Token</h4>
				<p class="font-mono text-sm bg-muted px-2 py-1 rounded">{course.token}</p>
			</div>
			<div>
				<h4 class="text-sm font-medium text-muted-foreground">Created</h4>
				<p class="text-sm">{new Date(course.createdAt).toLocaleDateString()}</p>
			</div>
			<div>
				<h4 class="text-sm font-medium text-muted-foreground">Last Updated</h4>
				<p class="text-sm">{new Date(course.updatedAt).toLocaleDateString()}</p>
			</div>
		</div>
	</CardContent>
</Card>