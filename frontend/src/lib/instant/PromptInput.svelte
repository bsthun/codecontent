<script lang="ts">
	import { createEventDispatcher } from 'svelte'
	import * as InputGroup from '$/lib/shadcn/components/input-group'
	import { ArrowUpIcon } from '@lucide/svelte'
	import { Separator } from '$/lib/shadcn/components/separator'

	export type Props = {
		value?: string
		placeholder?: string
		disabled?: boolean
		class?: string
		onSubmit?: (value: string) => void
	}

	let {
		value = $bindable(''),
		placeholder = 'Enter your prompt...',
		disabled = false,
		class: className,
		onSubmit
	}: Props = $props()

	const dispatch = createEventDispatcher<{
		submit: string
	}>()

	const handleSubmit = () => {
		if (value.trim() && !disabled) {
			onSubmit?.(value)
			dispatch('submit', value)

			// Mock action as alert
			alert(`Prompt sent: ${value}`)

			// Clear input after submission
			value = ''
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
				disabled={disabled || !value.trim()}
				onclick={handleSubmit}
			>
				<ArrowUpIcon />
				<span class="sr-only">Send</span>
			</InputGroup.Button>
		</InputGroup.Addon>
	</InputGroup.Root>
</div>