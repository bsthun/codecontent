<script lang="ts">
	import { navigate, useLocation } from 'svelte-navigator'
	import { getContext, onMount } from 'svelte'
	import type { Writable } from 'svelte/store'
	import type { Setup } from '$/util/type/setup.ts'
	import { toast } from 'svelte-sonner'
	import { backend } from '$/util/backend.ts'
	import Container from '$/component/layout/Container.svelte'
	import Loading from '$/component/interact/Loading.svelte'
	import { Card, CardContent, CardHeader, CardTitle } from '$/lib/shadcn/components/card'

	const location = useLocation()
	const params = new URLSearchParams($location.search)
	const code = params.get('code')
	const setup = getContext<Writable<Setup>>('setup')

	const processCallback = () => {
		if (!code) {
			setTimeout(() => {
				navigate('/entry/login')
			}, 3000)
			return
		}

		backend.public
			.loginCallback({
				code: code,
			})
			.then(() => {
				toast.success('Successfully logged in')
				$setup.reload().then(() => {
					navigate('/')
				})
			})
			.catch((err) => {
				if (err.response?.data) {
					toast.error(err.response.data.message, {
						description: err.response.data.error,
					})
				} else {
					toast.error(err.message)
				}

				setTimeout(() => {
					navigate('/entry/login')
				}, 3000)
			})
	}

	onMount(() => {
		processCallback()
	})
</script>

<Container class="flex min-h-dvh items-center justify-center">
	<Card class="mx-4 w-full max-w-md">
		<CardHeader class="text-center">
			<CardTitle class="justify-center">Processing Login</CardTitle>
		</CardHeader>
		<CardContent class="space-y-6 text-center">
			<p class="text-muted-foreground">Please wait while we verify your credentials</p>
			<div class="mb-4 flex justify-center">
				<Loading size="lg" />
			</div>
		</CardContent>
	</Card>
</Container>
