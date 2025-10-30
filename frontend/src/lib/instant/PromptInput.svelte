<script lang="ts">
	import * as InputGroup from '$/lib/shadcn/components/input-group'
	import { ArrowUpIcon } from '@lucide/svelte'
	import { Separator } from '$/lib/shadcn/components/separator'
	import { backend, catcher } from '$/util/backend'
	import { useNavigate } from 'svelte-navigator'
	import Loading from '$/component/interact/Loading.svelte'

	export type Props = {
		value?: string
		placeholder?: string
		disabled?: boolean
		class?: string
		courseId: number
	}

	let {
		value = $bindable(''),
		placeholder = 'Enter your prompt...',
		disabled = false,
		class: className,
		courseId
	}: Props = $props()

	const navigate = useNavigate()
	let loading = $state(false)

	const handleSubmit = () => {
		if (value.trim() && !disabled && !loading) {
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
	<InputGroup.Root>
		<InputGroup.Input
			bind:value
			{placeholder}
			{disabled}
			onkeydown={handleKeydown}
		/>
		<InputGroup.Addon align="inline-end">
			<Separator orientation="vertical" class="!h-4" />
			<InputGroup.Button
				variant="default"
				class="rounded-full"
				size="icon-xs"
				disabled={disabled || !value.trim() || loading}
				onclick={handleSubmit}
			>
				{#if loading}
					<Loading container={false} size="sm" />
				{:else}
					<ArrowUpIcon />
				{/if}
				<span class="sr-only">Send</span>
			</InputGroup.Button>
		</InputGroup.Addon>
	</InputGroup.Root>
</div>