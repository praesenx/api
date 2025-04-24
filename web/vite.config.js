import path from 'node:path';
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import tailwindcss from '@tailwindcss/vite';

export default defineConfig({
	define: {
		'process.env.VITE_API_URL': JSON.stringify(process.env.VITE_API_URL),
	},
	plugins: [vue(), tailwindcss()],
	resolve: {
		alias: [
			{
				find: /^~.+/,
				replacement: (val) => {
					return val.replace(/^~/, '');
				},
			},
			{
				find: '@',
				// eslint-disable-next-line no-undef
				replacement: path.resolve(__dirname, './src'),
			},
			{
				find: '@css',
				// eslint-disable-next-line no-undef
				replacement: path.resolve(__dirname, './src/css'),
			},
			{
				find: '@pages',
				// eslint-disable-next-line no-undef
				replacement: path.resolve(__dirname, './src/pages'),
			},
			{
				find: '@partials',
				// eslint-disable-next-line no-undef
				replacement: path.resolve(__dirname, './src/partials'),
			},
			{
				find: '@images',
				// eslint-disable-next-line no-undef
				replacement: path.resolve(__dirname, './src/images'),
			},
			{
				find: '@public',
				// eslint-disable-next-line no-undef
				replacement: path.resolve(__dirname, './src/public'),
			},
			{
				find: '@fonts',
				// eslint-disable-next-line no-undef
				replacement: path.resolve(__dirname, './src/fonts'),
			},
			{
				find: '@response',
				// eslint-disable-next-line no-undef
				replacement: path.resolve(__dirname, './src/response'),
			},
		],
	},
});
