<script lang="ts">
	import { SearchIcon, XIcon } from 'lucide-svelte'
	import { createEventDispatcher } from 'svelte'

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

<div class="form-control w-full">
	<div class="input-group relative">
		<div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none">
			<SearchIcon class="w-5 h-5 text-base-content/40" />
		</div>
		<input
			type="text"
			{placeholder}
			class="input input-bordered w-full pl-10 pr-10"
			bind:value
			on:input={handleInput}
		/>
		{#if value}
			<button
				class="absolute right-3 top-1/2 -translate-y-1/2 btn btn-ghost btn-xs btn-circle"
				on:click={clearSearch}
				aria-label="Clear search"
			>
				<XIcon class="w-4 h-4" />
			</button>
		{/if}
	</div>
</div>
