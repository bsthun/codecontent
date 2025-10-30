<script lang="ts">
	import { CalendarIcon, FileTextIcon, LayersIcon, ActivityIcon, ChevronRightIcon } from '@lucide/svelte'
	import { Link } from 'svelte-navigator'
	import { Card, CardHeader, CardTitle, CardDescription, CardContent } from '$/lib/shadcn/components/card'
	import { Button } from '$/lib/shadcn/components/button'
	import { formatDateTime } from '$/util/format'
	import type { PayloadContentInfo } from '$/util/backend/backend'
	import { CaseSensitiveIcon, WholeWordIcon } from 'lucide-svelte'

	export type Props = {
		content: PayloadContentInfo
		class?: string
	}

	let { class: className, content }: Props = $props()
</script>

<Link to={`/content/${content.id}/document`}>
	<Card class="{className} group cursor-pointer transition-all duration-200 hover:shadow-md hover:ring-2 hover:ring-primary/20">
		<CardHeader>
			<div class="flex items-start justify-between">
				<div class="flex-1 min-w-0">
					<CardTitle class="text-lg leading-tight group-hover:text-primary transition-colors duration-200">
						{content.title}
					</CardTitle>
				</div>
				<Button
					variant="ghost"
					size="icon"
					class="h-8 w-8 flex-shrink-0 opacity-0 group-hover:opacity-100 transition-opacity duration-200"
				>
					<ChevronRightIcon class="h-4 w-4" />
				</Button>
			</div>
		</CardHeader>

		<CardContent class="pt-0">
			<div class="space-y-4">
				<div class="grid grid-cols-2 gap-3">
					<div class="flex items-center gap-2 rounded-lg bg-muted/30 p-3 transition-colors duration-200 group-hover:bg-muted/50">
						<LayersIcon class="h-4 w-4 text-blue-600" />
						<div class="min-w-0 flex-1">
							<p class="text-sm font-medium">{content.contentSectionCount}</p>
							<p class="text-muted-foreground text-xs">Sections</p>
						</div>
					</div>
					<div class="flex items-center gap-2 rounded-lg bg-muted/30 p-3 transition-colors duration-200 group-hover:bg-muted/50">
						<ActivityIcon class="h-4 w-4 text-green-600" />
						<div class="min-w-0 flex-1">
							<p class="text-sm font-medium">{content.contentLogCount}</p>
							<p class="text-muted-foreground text-xs">Logs</p>
						</div>
					</div>
				</div>

				<!-- Date Information -->
				<div class="space-y-2 border-t border-muted/30 pt-3">
					<div class="flex items-center justify-between text-sm">
						<span class="flex items-center gap-2 text-muted-foreground">
							<CaseSensitiveIcon class="h-3 w-3" />
							Word Count
						</span>
						<span class="text-xs">0</span>
					</div>
					<div class="flex items-center justify-between text-sm">
						<span class="flex items-center gap-2 text-muted-foreground">
							<FileTextIcon class="h-3 w-3" />
							Updated
						</span>
						<span class="text-xs">{formatDateTime(content.updatedAt)}</span>
					</div>
				</div>
			</div>
		</CardContent>
	</Card>
</Link>