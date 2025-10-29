<script lang="ts">
	import Container from '$/component/layout/Container.svelte'
	import { ArrowLeftIcon } from 'lucide-svelte'
	import { Link } from 'svelte-navigator'
	import { Card, CardContent, CardHeader, CardTitle } from '$/lib/shadcn/components/card'
	import { Button } from '$/lib/shadcn/components/button'
</script>

<Container class="px-6">
	<div class="flex min-h-[80vh] items-center justify-center">
		<Card class="w-full max-w-md">
			<CardHeader class="flex flex-col items-center space-y-4 text-center">
				<div class="bg-secondary mb-4 rounded-full px-8 py-6">
					<span class="text-foreground text-6xl font-bold">404</span>
				</div>
				<CardTitle class="text-2xl">Page Not Found</CardTitle>
			</CardHeader>
			<CardContent class="flex flex-col items-center space-y-6 text-center">
				<p class="text-muted-foreground">The page you are looking for doesn't exist.</p>
				<Link to="/">
					<Button class="gap-2">
						<ArrowLeftIcon class="h-4 w-4" />
						Back to Home
					</Button>
				</Link>
			</CardContent>
		</Card>
	</div>
</Container>
