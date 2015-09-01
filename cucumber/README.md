# Dredd Cross Language Hooks Test Suite

Language agnostic CLI test suite for boilerplating [Dredd hooks](https://github.com/apiaryio/dredd) handler in new language written in [Aruba](https://github.com/cucumber/aruba)

## Usage

0. Install:
```
npm install -g dredd
git clone https://github.com/apiaryio/dredd-hooks-template.git
cd dredd-hooks-template
bundle install
```

1. Open feature files in `./features/*.feature`

2. In both features files replace:
  - `{{mylanguage}}` by hooks handler command for you language
  - `{{myextension}}` by extension for your language

3. Implement code examples in your language in both features

4. Run the test suite

```
$ bundle exec cucumber
```

5. Make `bundle exec cucumber` part of your test suite and CI
