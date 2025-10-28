<script lang="ts">
	import Navbar from '$/component/navbar/Navbar.svelte'
	import { getContext } from 'svelte'
	import type { Writable } from 'svelte/store'
	import type { Setup } from '$/util/type/setup'
	import { useLocation } from 'svelte-navigator'

	const setup = getContext<Writable<Setup>>('setup')
	const location = useLocation()

	let navbar = $state(false)
	
	$effect(() => {
		navbar = !$location.pathname.startsWith('/admin') && !$location.pathname.startsWith('/project/')
	})
	
	$effect(() => {
		if (!$setup.profile.id) {
			window.location.href = '/entry/login/'
		}
	})
</script>

<div>
	{#if navbar}
		<Navbar />
	{/if}
	<div>
		<slot />
	</div>
</div>
