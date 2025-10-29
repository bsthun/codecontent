<script lang="ts">
	import { SearchIcon, XIcon } from 'lucide-svelte'
	import { createEventDispatcher } from 'svelte'
	import { Button } from '$/lib/shadcn/components/button'

	export let value: string = ''
	export let placeholder: string = 'Search courses...'

	const dispatch = createEventDispatcher()

	const handleInput = (e: Event) => {
		const target = e.target as HTMLInputElement
		value = target.value
		dispatch('search', value)
	}

	const clearSearch = () => {
		value = ''
		dispatch('search', value)
	}
</script>

<div class="relative">
	<SearchIcon class="text-muted-foreground absolute top-1/2 left-3 h-4 w-4 -translate-y-1/2" />
	<input
		type="text"
		{placeholder}
		class="border-input bg-background ring-offset-background placeholder:text-muted-foreground focus-visible:ring-ring flex h-10 w-full rounded-md border px-3 py-2 pr-9 pl-9 text-sm file:border-0 file:bg-transparent file:text-sm file:font-medium focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
		bind:value
		on:input={handleInput}
	/>
	{#if value}
		<Button
			variant="ghost"
			size="icon"
			class="absolute top-1/2 right-1 h-6 w-6 -translate-y-1/2"
			onclick={clearSearch}
			aria-label="Clear search"
		>
			<XIcon class="h-3 w-3" />
		</Button>
	{/if}
</div>
