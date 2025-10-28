<script lang="ts">
	import { Loader2Icon } from 'lucide-svelte'
	import { backend, catcher } from '$/util/backend.ts'
	import Container from '$/component/layout/Container.svelte'

	let loading = false

	const handleLogin = () => {
		loading = true
		backend.public
			.loginRedirect()
			.then((res) => {
				window.location.href = res.data.redirectUrl
			})
			.catch((err) => {
				catcher(err)
				loading = false
			})
	}
</script>

<Container class="flex min-h-dvh items-center justify-center">
	<div class="card mx-4 w-full max-w-md bg-base-100 shadow-xl">
		<div class="card-body">
			<h2 class="card-title text-base-content">Login</h2>
			<p class="text-base-content/70 mb-6">Sign in to your account</p>
			<div class="card-actions">
				<button class="btn btn-primary w-full" disabled={loading} on:click={handleLogin}>
					{#if loading}
						<Loader2Icon class="mr-2 h-4 w-4 animate-spin" />
					{/if}
					Continue with OAuth
				</button>
			</div>
		</div>
	</div>
</Container>
