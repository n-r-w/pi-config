---
name: bundled-html-parser
description: Use when extracting local assets, readable content, or a normal local page from self-unpacking HTML files that store resources in __bundler script blocks.
---

<bundled_html_parser>
  <overview>
    This skill handles HTML files that embed page resources inside JSON script blocks instead of exposing normal asset files.

    Supported markers:
    - `<script type="__bundler/manifest">`
    - `<script type="__bundler/template">`
    - optional `<script type="__bundler/ext_resources">`
  </overview>

  <quick_start>
    Create `unpack-bundled-html.js` next to the input HTML, then run:

    ```bash
    HTML=input.html OUT=unpacked node unpack-bundled-html.js
    open unpacked/index.html
    less unpacked/strings.txt
    ```

    If local browser loading fails, serve the output directory:

    ```bash
    cd unpacked
    python3 -m http.server 8000
    ```

    Open `http://localhost:8000`.
  </quick_start>

  <node_unpacker>
    Use this script as the default unpacker:

    ```js
    const fs = require('node:fs');
    const path = require('node:path');
    const zlib = require('node:zlib');

    const input = process.env.HTML || 'input.html';
    const outDir = process.env.OUT || 'unpacked';
    const keywords = (process.env.KEYWORDS || '')
      .split(',')
      .map((value) => value.trim().toLowerCase())
      .filter(Boolean);

    fs.mkdirSync(outDir, { recursive: true });
    fs.mkdirSync(path.join(outDir, 'assets'), { recursive: true });

    const html = fs.readFileSync(input, 'utf8');

    function escapeRegExp(value) {
      return value.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
    }

    function getBundlerScript(type, required = true) {
      const safeType = escapeRegExp(type);
      const re = new RegExp(
        `<script\\s+type=["']${safeType}["']\\s*>\\s*([\\s\\S]*?)\\s*</script>`,
        'i',
      );
      const match = html.match(re);
      if (!match && required) throw new Error(`Missing <script type="${type}">`);
      return match ? match[1] : null;
    }

    function extFromMime(mime) {
      const map = {
        'text/javascript': '.js',
        'application/javascript': '.js',
        'text/css': '.css',
        'text/html': '.html',
        'application/json': '.json',
        'image/svg+xml': '.svg',
        'image/png': '.png',
        'image/jpeg': '.jpg',
        'image/webp': '.webp',
        'font/woff2': '.woff2',
        'font/woff': '.woff',
      };
      return map[mime] || '.bin';
    }

    function decodeAsset(entry) {
      let bytes = Buffer.from(entry.data, 'base64');
      if (entry.compressed) bytes = zlib.gunzipSync(bytes);
      return bytes;
    }

    function isReadableString(value) {
      if (value.length <= 3) return false;
      if (keywords.length) {
        const lower = value.toLowerCase();
        return keywords.some((keyword) => lower.includes(keyword));
      }
      return /\p{L}/u.test(value);
    }

    function extractStrings(filePath) {
      const src = fs.readFileSync(filePath, 'utf8');
      const values = [];
      const re = /"([^"\\]*(?:\\.[^"\\]*)*)"|'([^'\\]*(?:\\.[^'\\]*)*)'/g;

      for (const match of src.matchAll(re)) {
        const raw = match[1] ?? match[2];
        const value = raw
          .replace(/\\n/g, ' ')
          .replace(/\\"/g, '"')
          .replace(/\\'/g, "'")
          .trim();

        if (isReadableString(value)) values.push(value);
      }

      return values;
    }

    const manifest = JSON.parse(getBundlerScript('__bundler/manifest'));
    let template = JSON.parse(getBundlerScript('__bundler/template'));
    const extResourcesRaw = getBundlerScript('__bundler/ext_resources', false);
    const extResources = extResourcesRaw ? JSON.parse(extResourcesRaw) : [];

    const names = {};
    const mimeCounts = {};
    let compressedCount = 0;
    const unknownMimes = new Set();

    for (const [uuid, entry] of Object.entries(manifest)) {
      const ext = extFromMime(entry.mime);
      if (ext === '.bin') unknownMimes.add(entry.mime || '<missing>');
      if (entry.compressed) compressedCount += 1;
      mimeCounts[entry.mime] = (mimeCounts[entry.mime] || 0) + 1;

      const name = `assets/${uuid}${ext}`;
      const bytes = decodeAsset(entry);
      fs.writeFileSync(path.join(outDir, name), bytes);
      names[uuid] = name;
    }

    for (const [uuid, name] of Object.entries(names)) {
      template = template.split(uuid).join(name);
    }

    template = template
      .replace(/\s+integrity="[^"]*"/gi, '')
      .replace(/\s+crossorigin="[^"]*"/gi, '');

    fs.writeFileSync(path.join(outDir, 'index.html'), template);

    const strings = [];
    for (const name of Object.values(names)) {
      if (!name.endsWith('.js')) continue;
      strings.push(...extractStrings(path.join(outDir, name)));
    }
    fs.writeFileSync(path.join(outDir, 'strings.txt'), [...new Set(strings)].join('\n'));

    function hasBareUuidReference(source, uuid) {
      return source.includes(`=\"${uuid}\"`)
        || source.includes(`='${uuid}'`)
        || source.includes(`url(${uuid})`)
        || source.includes(`url(\"${uuid}\")`)
        || source.includes(`url('${uuid}')`);
    }

    const remainingBareUuids = Object.keys(manifest).filter((uuid) => hasBareUuidReference(template, uuid));
    const resourceMap = {};
    for (const entry of extResources) {
      if (names[entry.uuid]) resourceMap[entry.id] = names[entry.uuid];
    }

    console.log(`Wrote ${path.join(outDir, 'index.html')}`);
    console.log(`Wrote ${Object.keys(names).length} assets to ${path.join(outDir, 'assets')}`);
    console.log(`Wrote ${path.join(outDir, 'strings.txt')}`);
    console.log('MIME counts:', JSON.stringify(mimeCounts));
    console.log('Compressed assets:', compressedCount);
    if (unknownMimes.size) console.log('Unknown MIME types:', [...unknownMimes].join(', '));
    if (remainingBareUuids.length) console.log('Remaining bare UUID references:', remainingBareUuids.join(', '));
    if (Object.keys(resourceMap).length) console.log('External resource map:', JSON.stringify(resourceMap));
    ```
  </node_unpacker>

  <detection_rules>
    1. MUST use this skill only when the HTML has both `__bundler/manifest` and `__bundler/template` blocks.
    2. MUST treat bundler script content as JSON data, not executable JavaScript.
    3. MUST stop and use generic HTML parsing when either required block is missing.
  </detection_rules>

  <manifest_rules>
    1. Each manifest entry SHOULD contain `data`, `mime`, and `compressed`.
    2. `data` MUST be decoded from base64.
    3. `compressed: true` MUST be decompressed with gzip.
    4. Each asset MUST be saved as `assets/<uuid><extension>`.
    5. Unknown MIME types MUST be saved as `.bin` and reported.
  </manifest_rules>

  <template_rules>
    1. Every manifest UUID in the template MUST be replaced with the saved local asset path.
    2. `integrity="..."` MUST be removed from local output.
    3. `crossorigin="..."` MUST be removed from local output.
    4. The rewritten template MUST be saved as `index.html`.
    5. Remaining bare manifest UUID references MUST be reported. UUIDs inside generated `assets/<uuid><extension>` paths are expected.
  </template_rules>

  <text_extraction_rules>
    1. Extract strings from unpacked `.js` files when the user asks for readable content, labels, scenarios, keywords, or business text.
    2. Regex extraction from JavaScript strings is a fast first pass, not a complete JavaScript parse.
    3. Report the limitation when template literals, dynamic string assembly, minification, or runtime rendering may hide text.
    4. For higher confidence, render the unpacked page in a browser and extract visible DOM text.
  </text_extraction_rules>

  <search_commands>
    Use these commands after unpacking:

    ```bash
    KEYWORDS="keyword one,keyword two" HTML=input.html OUT=unpacked node unpack-bundled-html.js
    grep -n "keyword" unpacked/strings.txt
    grep -Rni "keyword" unpacked/assets
    ls unpacked/assets/*.js
    ```
  </search_commands>

  <validation_rules>
    1. MUST report asset count by MIME type.
    2. MUST report compressed asset count.
    3. MUST verify that `index.html` exists and has non-zero size.
    4. MUST verify that every manifest entry produced one asset file.
    5. SHOULD report `ext_resources` mappings when `ext_resources` is non-empty.
  </validation_rules>

  <safety_rules>
    1. MUST NOT execute unpacked JavaScript unless the user explicitly asks and the environment is safe.
    2. MUST NOT claim extracted text is complete unless DOM-rendered extraction or a real JavaScript parser was used.
    3. MUST keep original HTML unchanged unless the user explicitly asks to modify it.
  </safety_rules>

  <stop_rules>
    1. Stop and report a blocker if required script blocks are missing.
    2. Stop and report a blocker if JSON parsing fails.
    3. Stop and report a blocker if gzip decompression fails for a compressed asset.
    4. Ask for clarification if the user asks for semantic business interpretation but only raw extraction has been completed.
  </stop_rules>
</bundled_html_parser>
