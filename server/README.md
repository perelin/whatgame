# README

## GP

https://www.xbox.com/en-us/xbox-game-pass/games

## Heroku

deployment over github main branch (in heoku deploy console)
heroku logs -a p2-whatgame-server -t
heroku config -a p2-whatgame-server

## Twitch

https://dev.twitch.tv/console/apps/ofpcxa76vok6ldo3x22kw8yv7cti70
secret h6ydteyhysmenr0ab97yx55ez7j3cp
clientid ofpcxa76vok6ldo3x22kw8yv7cti70

## Reddit

https://www.reddit.com/r/XboxGamePass/comments/jt214y/public_api_for_fetching_the_list_of_game_pass/
Here's what the site uses:
First, it fetches different game lists (the one like "play with touch" or indie selection etc)
This is for the list of all GP Console games: https://catalog.gamepass.com/sigls/v2?id=f6f1f99f-9b49-4ccd-b3bf-4d9767a77f5e&language=en-us&market=US
This is for PC: https://catalog.gamepass.com/sigls/v2?id=fdd9e2a7-0fee-49f6-ad69-4354098401ff&language=en-us&market=US
EA Play: https://catalog.gamepass.com/sigls/v2?id=fdd9e2a7-0fee-49f6-ad69-4354098401ff&language=en-us&market=US
Play without a controller: https://catalog.gamepass.com/sigls/v2?id=7d8e8d56-c02f-4711-afec-73a80d8e9261&language=en-us&market=US
And finally all games: https://catalog.gamepass.com/sigls/v2?id=29a81209-df6f-41fd-a528-2ae6b91f719c&language=en-us&market=US
Of course change language and market accordingly. Each list will contain a list of product IDs
Then you can call this API to get the actual data:
https://displaycatalog.mp.microsoft.com/v7.0/products?bigIds={ids}&market=US&languages=en-us&MS-CV=DGU1mcuYo0WMMp
where {ids} is a comma separated list of product IDs, like BQ1W1T1FC14W,C3KLDKZBHNCZ,BS6WJ2L56B10. I don't know what MS-CV is and it doesn't work if you don't send it.
Another option is https://catalog.gamepass.com/products?market=US&language=en-US&hydration=MobileDetailsForConsole which accepts POST requests with a body like this
{
"Products": [ "BQ1W1T1FC14W", "C3KLDKZBHNCZ", "BS6WJ2L56B10" ]
}
I haven't found any other valid values for hydration, but IMHO this one works good enough (this one is used by the web version of xcloud)
