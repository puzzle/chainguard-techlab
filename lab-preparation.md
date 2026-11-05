# Lab preparation

## TODO

* [x] Ela zu Org einladen
* [x] 1.2 nix mehr optional, da in dev org
  * so belassen
* [ ] 2. Auth ausarbeiten (Ela)
* [ ] TODOs anschauen
* [ ] SBOMs 2x?
* [x] rename additional to advanced
* [x] 3.4 moove to advanced
* [ ] 4. helm charts ausarbeiten, wie finden, wie holen, images dazu? (Ela)
* [x] moove 5.1 to 5.2
* [ ] 5.1 wieso brauchen wir libs? wieso von Chinguard (SLSA oder Scanned / Ela)
* [ ] 5.2 ausarbeiten
* [ ] 6.2 hint to dfc, wie in 3.2
* [ ] 6.3 comparison ausarbeiten
* [ ] Präsis: Ela
* [ ] Ela Labs durchspielen mit spez org


## Ablauf / Agenda

* Intro Präsentation
  * Supply Chain Security
  * Chainguard Präsentation
* labs
  * 1.2: Account nicht anlegen, wird in 2 gemacht
* news / outro Präsentation

## Prepare

```bash
chainctl auth logout
docker logout cgr.dev
```


## invite users

https://console.chainguard.dev/org/chainguard.dev-partners/settings/users#/invite

Role: `registry.pull`

Copy invite link and send it to the user!

Create Role Binding for user:
* `libraries.java.pull`
* `libraries.java.pull_token_creator`

### user

* click on invite link
* enter email address of invite
* create account
* enter code from email
* 2fa
* any name for company
* join org (use invite code from invite)
* you should see `chainguard.dev-partners` org
