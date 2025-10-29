<script lang="ts">
	import { Edit2Icon, SaveIcon, XIcon } from '@lucide/svelte'
	import { Card, CardHeader, CardTitle, CardDescription, CardContent } from '$/lib/shadcn/components/card'
	import { Button } from '$/lib/shadcn/components/button'
	import { Textarea } from '$/lib/shadcn/components/textarea'

	export type Props = {
		promptInstruction: string
		loading?: boolean
		class?: string
		onSave?: (instruction: string) => void
	}

	let { class: className, promptInstruction, loading = false, onSave }: Props = $props()

	let editing = $state(false)
	let instruction = $state(promptInstruction)
	let tempInstruction = $state('')

	function startEdit() {
		tempInstruction = instruction
		editing = true
	}

	function cancelEdit() {
		editing = false
		tempInstruction = ''
	}

	function handleSave() {
		if (onSave) {
			onSave(tempInstruction)
		}
		editing = false
		instruction = tempInstruction
	}
</script>

<Card class={className}>
	<CardHeader>
		<div class="flex items-center justify-between">
			<div>
				<CardTitle>Prompt Instructions</CardTitle>
				<CardDescription>AI instructions for course content generation</CardDescription>
			</div>
			{#if !editing}
				<Button variant="outline" size="sm" onclick={startEdit}>
					<Edit2Icon class="h-4 w-4" />
				</Button>
			{:else}
				<div class="flex gap-2">
					<Button variant="outline" size="sm" onclick={cancelEdit} disabled={loading}>
						<XIcon class="h-4 w-4" />
					</Button>
					<Button size="sm" onclick={handleSave} disabled={loading}>
						<SaveIcon class="h-4 w-4" />
					</Button>
				</div>
			{/if}
		</div>
	</CardHeader>
	<CardContent>
		{#if editing}
			<Textarea
				bind:value={tempInstruction}
				placeholder="Enter prompt instructions for AI..."
				rows={8}
				class="w-full"
			/>
		{:else}
			<div class="min-h-[200px] p-4 bg-muted/50 rounded-md">
				{#if instruction}
					<p class="whitespace-pre-wrap text-sm">{instruction}</p>
				{:else}
					<p class="text-muted-foreground text-sm italic">No prompt instructions set</p>
				{/if}
			</div>
		{/if}
	</CardContent>
</Card>