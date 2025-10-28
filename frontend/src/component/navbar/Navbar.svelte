<script lang="ts">
	import { Link, navigate } from 'svelte-navigator'
	import { getContext } from 'svelte'
	import type { Writable } from 'svelte/store'
	import type { Setup } from '$/util/type/setup'
	import { onMount } from 'svelte'

	let scrolled = false

	const setup = getContext<Writable<Setup>>('setup')

	onMount(() => {
		const handleScroll = () => {
			scrolled = window.scrollY > 20
		}

		window.addEventListener('scroll', handleScroll)

		return () => {
			window.removeEventListener('scroll', handleScroll)
		}
	})
</script>

<div class="navbar fixed top-0 start-0 z-20 w-full bg-base-100 shadow-sm transition-all duration-300 max-lg:h-12" class:scrolled>
	<div class="navbar-start">
		<Link to="/" class="btn btn-ghost text-lg font-normal">
			Code Content Manager
		</Link>
	</div>
	<div class="navbar-center">
		<!-- Navigation items can be added here -->
	</div>
	<div class="navbar-end">
		<!-- Right side items can be added here -->
	</div>
</div>

<style lang="postcss">
	@reference '$/style/tailwind.css';

	.navbar {
		@apply border-b border-base-300 backdrop-blur-sm;
	}

	.navbar.scrolled {
		@apply shadow-md bg-base-200/80 backdrop-blur-md;
	}
</style>
