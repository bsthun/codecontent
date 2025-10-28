<script lang="ts">
	import Container from '$/component/layout/Container.svelte'
	import { ArrowLeftIcon } from 'lucide-svelte'
	import { Link } from 'svelte-navigator'
</script>

<Container class="px-6">
	<div class="flex min-h-[80vh] items-center justify-center">
		<div class="card w-full max-w-md bg-base-100 shadow-xl">
			<div class="card-body flex flex-col items-center text-center">
				<div class="bg-base-200 mb-4 rounded-full px-8 py-6">
					<span class="text-6xl font-bold text-base-content">404</span>
				</div>
				<h1 class="text-base-content card-title justify-center text-2xl">Page Not Found</h1>
				<p class="text-base-content/70 mb-6">The page you are looking for doesn't exist.</p>
				<div class="card-actions justify-center">
					<Link to="/">
						<button class="btn btn-primary gap-2">
							<ArrowLeftIcon class="h-4 w-4" />
							Back to Home
						</button>
					</Link>
				</div>
			</div>
		</div>
	</div>
</Container>
