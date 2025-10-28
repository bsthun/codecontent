<script lang="ts">
	export let title: string
	export let description: string | undefined = undefined
</script>

<div class="space-y-2">
	<h1 class="text-3xl font-bold tracking-tight text-base-content">{title}</h1>
	{#if description}
		<p class="text-base-content/70">{description}</p>
	{/if}
</div>
