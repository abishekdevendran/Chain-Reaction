<template>
		<div>
			<!-- The NuxtUI component will handle the background color for dark/light mode -->
			<UContainer
				class="flex flex-col items-center justify-center min-h-screen">
				<!-- Test for @nuxt/fonts -->
				<h1 class="text-4xl font-bold mb-4">Chain Reaction Game</h1>

				<!-- Test for @nuxt/icon -->
				<p
					class="text-lg text-gray-500 dark:text-gray-400 flex items-center gap-2 mb-8">
					<Icon name="i-heroicons-bolt" class="w-6 h-6 text-primary" />
					<span>Backend: Go | Frontend: Nuxt 4</span>
				</p>

				<!-- Test for @nuxt/ui -->
				<div class="flex gap-4">
					<UButton label="Login" size="xl" />
					<UButton label="Register" size="xl" color="white" variant="solid" />
				</div>
			</UContainer>
		</div>
</template>

<script setup lang="ts">
// Using Nuxt's composable to set the page title and metadata
useHead({
	title: 'Chain Reaction Game',
	meta: [
		{
			name: 'description',
			content: 'A multiplayer Chain Reaction game built with Go and Nuxt.'
		}
	]
});
</script>
