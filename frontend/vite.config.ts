import { sveltekit } from '@sveltejs/kit/vite';
// @ts-ignore
import path from 'path';
import { defineConfig } from 'vitest/config';

export default defineConfig({
	plugins: [sveltekit()],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	},
	server: {
		fs: {
			allow: [path.resolve('bindings/main')]
		}
	}
});
