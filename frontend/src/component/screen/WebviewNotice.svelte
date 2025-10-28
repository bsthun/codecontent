<script lang="ts">
	import Container from '$/component/layout/Container.svelte'
	import { AlertTriangleIcon, EllipsisVerticalIcon } from 'lucide-svelte'
</script>

<Container class="px-6">
	<div class="flex min-h-[80vh] items-center justify-center">
		<div class="card w-full max-w-md bg-base-100 shadow-xl">
			<div class="card-body flex flex-col items-center text-center">
				<div class="mb-4 rounded-full bg-warning/20 p-3">
					<AlertTriangleIcon class="h-10 w-10 text-warning" />
				</div>
				<h1 class="text-base-content card-title justify-center text-2xl mb-2">Webview is not supported</h1>
				<p class="text-base-content/70 mb-6">
					Please open this website in the native browser. You can click on the
					<EllipsisVerticalIcon
						class="inline animate-pulse border border-base-content p-0.5"
						size="1.2em"
					/>
					button and choose to open in browser.
				</p>
			</div>
		</div>
	</div>
</Container>
