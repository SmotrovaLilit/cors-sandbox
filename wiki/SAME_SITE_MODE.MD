## Cookie same site mode

Behavior in chromium version 107.0.5304.110.

"+" - means that a browser will send the cookie

"-" - means that a browser will not send the cookie

Main domain is different for sites, for example http://origin.cors.com
and http://allowing.cors-another.com

| Action/State                                   | SameSite=None  | Without sameSite | SameSite=Lax    | SameSite=strict |
|------------------------------------------------|----------------|------------------|-----------------|-----------------|
| Follow cross-site link                         | +              | +                | +               | -               |
| Submit form with action to the cross site link | GET: +, POST:+ | GET: +, POST:+   | GET: +, POST: - | -               |
| Follow cross-site redirect                     | +              | +                | +               | -               |
| Fetch (credentials=include, mode='no-cors')    | +              | -                | -               | -               |

## 💡Notes:
- Some browsers will block a cookie with SameSite=None if the Secure flag is not also set.
- We need to set same site mode cookie, because default value is different in browsers and in chromium behaviour without same site is different with other modes. 