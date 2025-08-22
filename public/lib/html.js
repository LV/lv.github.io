// Refined from https://plainvanillaweb.com/lib/html.js

class Html extends String { }


/**
 * Marks a string as `HTML` to not encode it
 *
 * Normally, strings are entity-encoded (e.g. `&lt;` instead of `<`)
 *   before being inserted into HTML
 *
 * The following wraps the string into the `Html` class to skip encoding
 *
 *
 * @param {string} str
 * @returns {Html}
 */
export const htmlRaw = str => new Html(str);


/**
 * Escapes special characters for safe insertion into HTML
 *
 *
 * @param {*} value
 * @returns {Html}
 */
export const htmlEncode = (value) => {
  // Avoid double-encoding if value is already `Html`
  if (value instanceof Html) return value;

  /** @type {Record<string, string>} */
  const map = {
    '&': '&amp;',
    '<': '&lt;',
    '>': '&gt;',
    "'": '&#39;',
    '"': '&quot;'
  };

  return htmlRaw(
    String(value).replace(/[&<>'"]/g, (tag) => map[tag])
  );
};


/**
 * Template literal tag that safely builds HTML by auto-encoding user data
 *
 * Example usage:
 *   html`<p>Hello, ${userName}</p>`
 *
 * When you write the above, the function receives:
 *   strings = ['<p>Hello, ', '</p>']    // The literal HTML parts
 *   values  = [userName]                // The ${} interpolated values
 *
 * It then:
 *   1. HTML-encodes each user value (so "<script>" becomes "&lt;script&gt;")
 *   2. Weaves the encoded values back to the HTML template
 *   3. Returns the result as safe HTML
 *
 *
 * Example:
 *   const userName = '<script>alert("xss")</script>';
 *   html`<p>Hello, ${userName}</p>`
 *   // Result: '<p>Hello, &lt;script&gt;alert("xss")&lt;/script&gt;</p>'
 *
 *
 * @param {TemplateStringsArray} strings - The literal parts of the template
 * @param {...*} values - The interpolated ${} values that need encoding
 * @returns {Html} - Safe HTML string marked as already processed
 */
export const html = (strings, ...values) => 
  htmlRaw(
    String.raw(
      { raw: strings },
      // Transform each user value into a safe string before inserting
      ...values.map(v => String(htmlEncode(v)))
    )
  );

// `String.raw` does the following:
// String.raw({ raw: ['<p>Hello, ', '</p>'] }, 'SAFE_USER_DATA')
// Result: '<p>Hello, SAFE_USER_DATA</p>'
//
// Equivalently: strings[0] + values[0] + strings[1] + values[1] + ...
//   but preserves literal backslashes in the HTML templates
