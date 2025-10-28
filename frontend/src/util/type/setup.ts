export type Setup = {
	profile: {
		id?: string
		name?: string
		email?: string
		pictureUrl?: string
		isAdmin?: boolean
	}
	initialized: boolean
	reload: () => Promise<void>
}
