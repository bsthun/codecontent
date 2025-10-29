<script lang="ts">
	import { LoaderCircleIcon } from '@lucide/svelte'

	export type Props = {
		container?: boolean
		class?: string
		size?: 'sm' | 'md' | 'lg'
	}

	let { container = false, class: className = '', size = 'md' }: Props = $props()

	const getSizeClasses = () => {
		switch (size) {
			case 'sm':
				return 'h-4 w-4'
			case 'lg':
				return 'h-8 w-8'
			default:
				return 'h-6 w-6'
		}
	}
</script>

{#if container}
	<div class="flex min-h-screen items-center justify-center {className}">
		<LoaderCircleIcon class="animate-spin {getSizeClasses()}" />
	</div>
{:else}
	<LoaderCircleIcon class="animate-spin {getSizeClasses()} {className}" />
{/if}
