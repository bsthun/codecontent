<script lang="ts">
	import { Link } from 'svelte-navigator'
	import { onMount } from 'svelte'
	import { Button } from '$/lib/shadcn/components/button'

	let scrolled = false

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

<nav
	class="bg-background border-border fixed start-0 top-0 z-20 w-full border-b shadow-sm backdrop-blur-sm transition-all duration-300 max-lg:h-12 {scrolled
		? 'bg-secondary/80 shadow-md backdrop-blur-md'
		: ''}"
>
	<div class="flex h-16 items-center justify-between px-6 max-lg:px-4">
		<div class="flex items-center">
			<Link to="/">
				<Button variant="ghost" class="h-auto px-2 py-1 text-lg font-normal">Code Content Manager</Button>
			</Link>
		</div>
		<div class="flex flex-1 justify-center">
			<!-- Navigation items can be added here -->
		</div>
		<div class="flex items-center">
			<!-- Right side items can be added here -->
		</div>
	</div>
</nav>
