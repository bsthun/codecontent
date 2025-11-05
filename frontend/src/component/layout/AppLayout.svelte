<script lang="ts">
	import Navbar from '$/component/navbar/Navbar.svelte'
	import { getContext } from 'svelte'
	import type { Writable } from 'svelte/store'
	import type { Setup } from '$/util/type/setup'
	import { useLocation, useNavigate } from 'svelte-navigator'
	import type { Snippet } from 'svelte'

	export type Props = {
		children?: Snippet
	}

	let { children }: Props = $props()

	const setup = getContext<Writable<Setup>>('setup')
	const location = useLocation()
	const navigate = useNavigate()

	let navbar = $state(false)

	$effect(() => {
		navbar = !$location.pathname.startsWith('/admin') && !$location.pathname.startsWith('/project/')
	})

	$effect(() => {
		if (!$setup.profile.id) {
			navigate('/entry/login/')
		}
	})
</script>

<div>
	{#if navbar}
		<Navbar />
	{/if}
	<div>
		{@render children?.()}
	</div>
</div>
