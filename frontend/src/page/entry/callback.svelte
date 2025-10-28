<script lang="ts">
	import { navigate, useLocation } from 'svelte-navigator'
	import { getContext, onMount } from 'svelte'
	import type { Writable } from 'svelte/store'
	import type { Setup } from '$/util/type/setup.ts'
	import { toast } from 'svelte-sonner'
	import { backend } from '$/util/backend.ts'
	import Container from '$/component/layout/Container.svelte'

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
	<div class="card mx-4 w-full max-w-md bg-base-100 shadow-xl">
		<div class="card-body text-center">
			<h2 class="card-title text-base-content justify-center">
				Processing Login
			</h2>
			<p class="text-base-content/70 mb-6">
				Please wait while we verify your credentials
			</p>
			<div class="mb-4 flex justify-center">
				<div class="loading loading-spinner loading-lg text-base-content"></div>
			</div>
		</div>
	</div>
</Container>
