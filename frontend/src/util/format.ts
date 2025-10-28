export const formatDate = (dateString: string | undefined): string => {
	const date = new Date(dateString!)
	return date.toLocaleDateString('en-GB', {
		year: 'numeric',
		month: 'short',
		day: 'numeric',
	})
}

export const formatDateTime = (dateString: string): string => {
	const date = new Date(dateString)
	const formatted = date.toLocaleDateString('en-GB', {
		day: 'numeric',
		month: 'short',
		year: 'numeric',
		hour: 'numeric',
		minute: '2-digit',
		hour12: true,
	})
	return formatted.replace(/am|pm/gi, (match) => match.toUpperCase())
}

export const formatDuration = (durationMs: number): string => {
	const totalSeconds = Math.floor(durationMs / 1000)
	const minutes = Math.floor(totalSeconds / 60)
	const seconds = totalSeconds % 60
	return `${minutes.toString().padStart(1, '0')}:${seconds.toString().padStart(2, '0')}`
}
