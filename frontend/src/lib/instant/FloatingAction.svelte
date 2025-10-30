<script lang="ts">
	import { fly } from 'svelte/transition'
	import type { Snippet } from 'svelte'

	export type Props = {
		visible?: boolean
		class?: string
		children: Snippet
	}

	let { visible = true, class: className, children }: Props = $props()
</script>

{#if visible}
	<div
		class="fixed bottom-6 left-1/2 z-50 -translate-x-1/2 transform"
		transition:fly={{ y: 50, duration: 200 }}
		class:className
	>
		<div class="rounded-lg border bg-background p-4 shadow-lg backdrop-blur-sm">
			{@render children()}
		</div>
	</div>
{/if}