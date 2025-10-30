<script lang="ts">
	import { CalendarIcon, FileTextIcon, LayersIcon, ActivityIcon, ChevronRightIcon } from '@lucide/svelte'
	import { Card, CardHeader, CardTitle, CardDescription, CardContent } from '$/lib/shadcn/components/card'
	import { Button } from '$/lib/shadcn/components/button'
	import { formatDateTime } from '$/util/format'
	import type { PayloadContentInfo } from '$/util/backend/backend'

	export type Props = {
		content: PayloadContentInfo
		class?: string
	}

	let { class: className, content }: Props = $props()

	const handleCardClick = () => {
		// Navigate to content detail page - adjust route as needed
		window.location.href = `/course/content/${content.id}`
	}

	const stats = $derived([
		{
			icon: LayersIcon,
			label: 'Sections',
			value: content.contentSectionCount,
			color: 'text-blue-600'
		},
		{
			icon: ActivityIcon,
			label: 'Logs',
			value: content.contentLogCount,
			color: 'text-green-600'
		}
	])
</script>

<Card class="{className} group cursor-pointer transition-all duration-200 hover:shadow-md hover:ring-2 hover:ring-primary/20">
	<CardHeader class="pb-3">
		<div class="flex items-start justify-between gap-4">
			<div class="flex-1 min-w-0">
				<CardTitle class="text-lg leading-tight group-hover:text-primary transition-colors duration-200">
					{content.title}
				</CardTitle>
				<CardDescription class="mt-1 text-sm">
					Content ID: {content.id}
				</CardDescription>
			</div>
			<Button
				variant="ghost"
				size="icon"
				class="h-8 w-8 flex-shrink-0 opacity-0 group-hover:opacity-100 transition-opacity duration-200"
				onclick={handleCardClick}
			>
				<ChevronRightIcon class="h-4 w-4" />
			</Button>
		</div>
	</CardHeader>

	<CardContent class="pt-0">
		<div class="space-y-4">
			<!-- Stats Grid -->
			<div class="grid grid-cols-2 gap-3">
				{#each stats as stat}
					<div class="flex items-center gap-2 rounded-lg bg-muted/30 p-3 transition-colors duration-200 group-hover:bg-muted/50">
						<stat.icon class="h-4 w-4 {stat.color}" />
						<div class="min-w-0 flex-1">
							<p class="text-sm font-medium">{stat.value}</p>
							<p class="text-muted-foreground text-xs">{stat.label}</p>
						</div>
					</div>
				{/each}
			</div>

			<!-- Date Information -->
			<div class="space-y-2 border-t border-muted/30 pt-3">
				<div class="flex items-center justify-between text-sm">
					<span class="flex items-center gap-2 text-muted-foreground">
						<CalendarIcon class="h-3 w-3" />
						Created
					</span>
					<span class="text-xs">{formatDateTime(content.createdAt)}</span>
				</div>
				<div class="flex items-center justify-between text-sm">
					<span class="flex items-center gap-2 text-muted-foreground">
						<FileTextIcon class="h-3 w-3" />
						Updated
					</span>
					<span class="text-xs">{formatDateTime(content.updatedAt)}</span>
				</div>
			</div>

			<!-- Action Button -->
			<Button
				variant="outline"
				size="sm"
				class="w-full gap-2 transition-colors duration-200 group-hover:bg-primary group-hover:text-primary-foreground"
				onclick={handleCardClick}
			>
				<FileTextIcon class="h-4 w-4" />
				View Content
			</Button>
		</div>
	</CardContent>
</Card>