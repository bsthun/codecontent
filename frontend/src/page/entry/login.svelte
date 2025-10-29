<script lang="ts">
	import { backend, catcher } from '$/util/backend.ts'
	import Container from '$/component/layout/Container.svelte'
	import Loading from '$/component/interact/Loading.svelte'
	import { Card, CardContent, CardHeader, CardTitle } from '$/lib/shadcn/components/card'
	import { Button } from '$/lib/shadcn/components/button'

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
	<Card class="mx-4 w-full max-w-md">
		<CardHeader>
			<CardTitle>Login</CardTitle>
		</CardHeader>
		<CardContent class="space-y-6">
			<p class="text-muted-foreground">Sign in to your account</p>
			<Button class="w-full" disabled={loading} onclick={handleLogin}>
				{#if loading}
					<Loading size="sm" class="mr-2" />
				{/if}
				Continue with OAuth
			</Button>
		</CardContent>
	</Card>
</Container>
