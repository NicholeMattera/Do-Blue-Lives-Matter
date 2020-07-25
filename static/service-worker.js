
const CACHE_NAME = 'offline';
const SMATTER_OFFLINE_URL = '/smatter/offline.html';
const SREALLYMATTER_OFFLINE_URL = '/sreallymatter/offline.html';

self.addEventListener('install', (event) => {
    event.waitUntil((async () => {
        const cache = await caches.open(CACHE_NAME);
        await cache.add(new Request(SMATTER_OFFLINE_URL, {cache: 'reload'}));
        await cache.add(new Request(SREALLYMATTER_OFFLINE_URL, {cache: 'reload'}));
    })());
});

self.addEventListener('activate', (event) => {
    event.waitUntil((async () => {
        if ('navigationPreload' in self.registration) {
            await self.registration.navigationPreload.enable();
        }
    })());

    self.clients.claim();
});

self.addEventListener('fetch', (event) => {
    if (event.request.mode === 'navigate') {
        event.respondWith((async () => {
            try {
                const preloadResponse = await event.preloadResponse;
                return (preloadResponse) ? preloadResponse : await fetch(event.request);
            } catch (error) {
                console.log(event);
                const cache = await caches.open(CACHE_NAME);
                return await cache.match((event.request.url.endsWith('/smatter')) ? SMATTER_OFFLINE_URL : SREALLYMATTER_OFFLINE_URL);
            }
        })());
    }
});