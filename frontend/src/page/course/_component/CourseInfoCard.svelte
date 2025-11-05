<script lang="ts">
	import { Edit2Icon, Trash2Icon, CalendarIcon, ClockIcon, CopyIcon, CheckIcon, LinkIcon } from '@lucide/svelte'
	import { Card, CardHeader, CardTitle, CardDescription, CardContent } from '$/lib/shadcn/components/card'
	import { Button } from '$/lib/shadcn/components/button'
	import type { PayloadCourse } from '$/util/backend/backend'
	import { formatDateTime } from '$/util/format'
	import { catcher } from '$/util/backend.ts'

	export type Props = {
		course: PayloadCourse
		class?: string
		onEdit?: () => void
		onDelete?: () => void
	}

	let { class: className, course, onEdit, onDelete }: Props = $props()

	let copied = $state(false)
	let copyTimeout: ReturnType<typeof setTimeout>

	const inviteUrl = $derived(`${window.location.origin}/course/invite?courseId=${course.id}&token=${course.token}`)

	const copyToClipboard = () => {
		navigator.clipboard.writeText(inviteUrl)
			.then(() => {
				copied = true

				if (copyTimeout) {
					clearTimeout(copyTimeout)
				}
				copyTimeout = setTimeout(() => {
					copied = false
				}, 3000)
			})
			.catch((err) => {
				catcher(err)
			})
	}
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
		<div class="space-y-6">
			<div class="space-y-3">
				<div class="flex items-center gap-2">
					<LinkIcon class="h-4 w-4 text-muted-foreground" />
					<span class="text-sm font-medium text-muted-foreground">Invite Link</span>
				</div>
				<div class="group relative">
					<div
						role="button"
						tabindex="0"
						class="flex items-center gap-3 bg-gradient-to-r from-muted/50 to-muted/30 border border-muted/30 rounded-lg p-3 cursor-pointer transition-all duration-200 hover:border-primary/30 hover:shadow-sm focus:outline-none focus:ring-2 focus:ring-primary/50"
						onclick={copyToClipboard}
						onkeydown={(e) => e.key === 'Enter' && copyToClipboard()}
					>
						<div class="flex-1 min-w-0">
							<div class="font-mono text-sm text-foreground truncate">
								/course/invite?courseId={course.id}&token={course.token}
							</div>
							<div class="text-xs text-muted-foreground mt-1">
								{window.location.origin}
							</div>
						</div>
						<div class="flex-shrink-0">
							<Button
								variant="ghost"
								size="sm"
								class="h-8 w-8 p-0 transition-all duration-200 hover:bg-primary/10"
								onclick={copyToClipboard}
							>
								{#if copied}
									<CheckIcon class="h-4 w-4 text-green-600 animate-in fade-in duration-300" />
								{:else}
									<CopyIcon class="h-4 w-4 text-muted-foreground group-hover:text-foreground transition-colors duration-200" />
								{/if}
							</Button>
						</div>
					</div>
					{#if copied}
						<div class="absolute -top-8 right-0 bg-green-600 text-white text-xs px-2 py-1 rounded-md animate-in fade-in slide-in-from-bottom-2 duration-300">
							Copied!
						</div>
					{/if}
				</div>
			</div>

			<!-- Date Information -->
			<div class="space-y-3 border-t border-muted/30 pt-4">
				<div class="flex items-center justify-between">
					<span class="flex items-center gap-2 text-muted-foreground text-sm">
						<CalendarIcon class="h-4 w-4" />
						Created
					</span>
					<span class="text-sm font-medium">{formatDateTime(course.createdAt)}</span>
				</div>
				<div class="flex items-center justify-between">
					<span class="flex items-center gap-2 text-muted-foreground text-sm">
						<ClockIcon class="h-4 w-4" />
						Updated
					</span>
					<span class="text-sm font-medium">{formatDateTime(course.updatedAt)}</span>
				</div>
			</div>
		</div>
	</CardContent>
</Card>