<script lang="ts">
	import { ArrowUpIcon } from '@lucide/svelte'
	import { Separator } from '$/lib/shadcn/components/separator'
	import { backend, catcher } from '$/util/backend'
	import { useNavigate } from 'svelte-navigator'
	import Loading from '$/component/interact/Loading.svelte'
	import { InputGroup, InputGroupAddon, InputGroupButton, InputGroupInput } from '$/lib/shadcn/components/input-group'

	export type Props = {
		value?: string
		placeholder?: string
		class?: string
		courseId: number
	}

	let {
		value = $bindable(''),
		placeholder = 'Enter your prompt...',
		class: className,
		courseId
	}: Props = $props()

	const navigate = useNavigate()
	let loading = $state(false)

	const handleSubmit = () => {
		if (value.trim() && !loading) {
			loading = true

			backend.content
				.contentCreate({
					courseId,
					prompt: value.trim()
				})
				.then((response) => {
					const contentId = response.data.content.id
					navigate(`/content/${contentId}/document`)
				})
				.catch((err) => {
					catcher(err)
				})
				.finally(() => {
					loading = false
				})
		}
	}

	const handleKeydown = (event: KeyboardEvent) => {
		if (event.key === 'Enter' && !event.shiftKey) {
			event.preventDefault()
			handleSubmit()
		}
	}
</script>

<div class={className}>
	<InputGroup>
		<InputGroupInput
			bind:value
			{placeholder}
			disabled={loading}
			onkeydown={handleKeydown}
		/>
		<InputGroupAddon align="inline-end">
			<Separator orientation="vertical" class="!h-4" />
			<InputGroupButton
				variant="default"
				class="rounded-full"
				size="icon-xs"
				disabled={!value.trim() || loading}
				onclick={handleSubmit}
			>
				{#if loading}
					<Loading container={false} size="sm" />
				{:else}
					<ArrowUpIcon />
				{/if}
				<span class="sr-only">Send</span>
			</InputGroupButton>
		</InputGroupAddon>
	</InputGroup>
</div>