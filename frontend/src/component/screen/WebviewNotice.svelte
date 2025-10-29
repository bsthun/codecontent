<script lang="ts">
	import Container from '$/component/layout/Container.svelte'
	import { AlertTriangleIcon, EllipsisVerticalIcon } from 'lucide-svelte'
	import { Card, CardContent, CardHeader, CardTitle } from '$/lib/shadcn/components/card'
</script>

<Container class="px-6">
	<div class="flex min-h-[80vh] items-center justify-center">
		<Card class="w-full max-w-md">
			<CardHeader class="flex flex-col items-center space-y-4 text-center">
				<div class="mb-4 rounded-full bg-yellow-100 p-3 dark:bg-yellow-900/20">
					<AlertTriangleIcon class="h-10 w-10 text-yellow-600 dark:text-yellow-400" />
				</div>
				<CardTitle class="mb-2 text-2xl">Webview is not supported</CardTitle>
			</CardHeader>
			<CardContent class="space-y-6 text-center">
				<p class="text-muted-foreground">
					Please open this website in the native browser. You can click on the
					<EllipsisVerticalIcon class="border-border inline animate-pulse border p-0.5" size="1.2em" />
					button and choose to open in browser.
				</p>
			</CardContent>
		</Card>
	</div>
</Container>
