<!DOCTYPE html>
<html lang="ko">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>냉장Go - 스마트 레시피 추천</title>
  <script src="https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"></script>
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link href="https://fonts.googleapis.com/css2?family=Noto+Sans+KR:wght@300;400;500;700;900&display=swap" rel="stylesheet">
  <style>
    :root {
      --green: #2bb673; --green-light: #e8fff3; --green-dark: #1a9057;
      --accent: #ff6b35; --bg: #f4f7f5; --card: #ffffff;
      --text: #1a1a2e; --muted: #6b7280;
    }
    * { box-sizing: border-box; margin: 0; padding: 0; }
    body { font-family: 'Noto Sans KR', sans-serif; background: var(--bg); color: var(--text); min-height: 100vh; }

    /* ── HEADER ── */
    header { background: #fff; border-bottom: 1px solid #e5e7eb; position: sticky; top: 0; z-index: 100; }
    .header-inner { max-width: 1200px; margin: 0 auto; padding: 0 24px; height: 64px; display: flex; align-items: center; justify-content: space-between; }
    .logo { font-size: 24px; font-weight: 900; color: var(--green); letter-spacing: -1px; cursor: pointer; }
    .logo span { color: var(--text); }
    nav a { font-size: 14px; font-weight: 500; color: #555; text-decoration: none; margin-left: 28px; transition: color .2s; cursor: pointer; }
    nav a:hover, nav a.active { color: var(--green); }
    .header-auth { display: flex; align-items: center; gap: 10px; }
    .login-btn { background: var(--green); color: #fff; border: none; border-radius: 999px; padding: 8px 20px; font-size: 13px; font-weight: 700; cursor: pointer; transition: background .2s; font-family: inherit; }
    .login-btn:hover { background: var(--green-dark); }
    .login-btn.outline { background: transparent; border: 2px solid var(--green); color: var(--green); }
    .login-btn.outline:hover { background: var(--green-light); }
    .user-pill { display: flex; align-items: center; gap: 8px; background: var(--green-light); border-radius: 999px; padding: 6px 14px 6px 8px; cursor: pointer; }
    .user-avatar { width: 28px; height: 28px; border-radius: 50%; background: var(--green); color: #fff; display: flex; align-items: center; justify-content: center; font-size: 12px; font-weight: 800; }
    .user-name { font-size: 13px; font-weight: 700; color: var(--green-dark); }
    .logout-btn { background: none; border: none; font-size: 12px; color: #aaa; cursor: pointer; font-family: inherit; transition: color .2s; }
    .logout-btn:hover { color: #ef4444; }

    /* ── AUTH MODAL ── */
    .auth-overlay { display: none; position: fixed; inset: 0; background: rgba(0,0,0,.5); z-index: 2000; align-items: center; justify-content: center; padding: 20px; backdrop-filter: blur(6px); }
    .auth-overlay.open { display: flex; }
    .auth-box { background: #fff; border-radius: 28px; width: 100%; max-width: 440px; padding: 40px; box-shadow: 0 40px 80px rgba(0,0,0,.2); animation: fadeIn .25s ease; }
    .auth-logo { font-size: 28px; font-weight: 900; color: var(--green); text-align: center; margin-bottom: 4px; }
    .auth-sub { font-size: 13px; color: var(--muted); text-align: center; margin-bottom: 28px; }
    .auth-tabs { display: flex; background: #f4f7f5; border-radius: 12px; padding: 4px; margin-bottom: 24px; gap: 4px; }
    .auth-tab { flex: 1; padding: 9px; border: none; border-radius: 9px; font-size: 14px; font-weight: 700; cursor: pointer; font-family: inherit; background: transparent; color: #999; transition: all .2s; }
    .auth-tab.active { background: #fff; color: var(--green); box-shadow: 0 2px 8px rgba(0,0,0,.08); }
    .auth-form { display: flex; flex-direction: column; gap: 12px; }
    .auth-form label { font-size: 12px; font-weight: 700; color: #888; letter-spacing: .4px; margin-bottom: -4px; }
    .auth-form input { border: 2px solid #e5e7eb; border-radius: 12px; padding: 12px 16px; font-size: 15px; font-family: inherit; outline: none; transition: border .2s; }
    .auth-form input:focus { border-color: var(--green); }
    .auth-submit { background: var(--green); color: #fff; border: none; border-radius: 12px; padding: 14px; font-size: 15px; font-weight: 700; cursor: pointer; font-family: inherit; margin-top: 4px; transition: background .2s; }
    .auth-submit:hover { background: var(--green-dark); }
    .auth-error { background: #fee2e2; color: #991b1b; border-radius: 10px; padding: 10px 14px; font-size: 13px; font-weight: 600; display: none; }
    .auth-error.show { display: block; }
    .auth-close { position: absolute; top: 16px; right: 16px; background: #f1f3f2; border: none; border-radius: 50%; width: 34px; height: 34px; font-size: 16px; cursor: pointer; display: flex; align-items: center; justify-content: center; }
    .auth-box { position: relative; }

    /* ── DB ADMIN PANEL ── */
    .db-overlay { display: none; position: fixed; inset: 0; background: rgba(0,0,0,.6); z-index: 3000; align-items: center; justify-content: center; padding: 20px; backdrop-filter: blur(4px); }
    .db-overlay.open { display: flex; }
    .db-panel { background: #1a1a2e; border-radius: 20px; width: 100%; max-width: 860px; max-height: 86vh; overflow: hidden; display: flex; flex-direction: column; box-shadow: 0 40px 80px rgba(0,0,0,.5); animation: fadeIn .2s ease; }
    .db-header { padding: 20px 24px; border-bottom: 1px solid #2d2d50; display: flex; align-items: center; justify-content: space-between; }
    .db-title { font-size: 16px; font-weight: 800; color: #7dd3fc; letter-spacing: .3px; }
    .db-tabs { display: flex; gap: 8px; }
    .db-tab { padding: 6px 16px; border-radius: 8px; border: 1px solid #2d2d50; font-size: 13px; font-weight: 700; cursor: pointer; font-family: inherit; color: #aaa; background: transparent; transition: all .2s; }
    .db-tab.active { background: #7dd3fc20; border-color: #7dd3fc50; color: #7dd3fc; }
    .db-close { background: #2d2d50; border: none; border-radius: 8px; padding: 6px 14px; color: #aaa; cursor: pointer; font-size: 13px; font-family: inherit; }
    .db-body { flex: 1; overflow-y: auto; padding: 20px 24px; }
    .db-table { width: 100%; border-collapse: collapse; font-size: 13px; }
    .db-table th { text-align: left; padding: 8px 12px; color: #7dd3fc; font-weight: 700; border-bottom: 1px solid #2d2d50; white-space: nowrap; }
    .db-table td { padding: 8px 12px; color: #c8d6e5; border-bottom: 1px solid #1e1e38; vertical-align: top; word-break: break-word; max-width: 200px; }
    .db-table tr:hover td { background: #ffffff08; }
    .db-badge { display: inline-block; padding: 2px 8px; border-radius: 999px; font-size: 11px; font-weight: 700; }
    .db-badge.green { background: #064e3b; color: #6ee7b7; }
    .db-badge.gray { background: #2d2d50; color: #9ca3af; }
    .db-empty { color: #555; font-size: 13px; padding: 20px 0; text-align: center; }
    .db-stat-row { display: flex; gap: 12px; margin-bottom: 16px; }
    .db-stat { background: #2d2d50; border-radius: 10px; padding: 12px 16px; flex: 1; }
    .db-stat-n { font-size: 22px; font-weight: 900; color: #7dd3fc; }
    .db-stat-l { font-size: 11px; color: #888; margin-top: 2px; }
    .sql-box { background: #0d0d1a; border-radius: 12px; padding: 16px; margin-bottom: 12px; }
    .sql-box textarea { width: 100%; background: transparent; border: none; outline: none; color: #a3e635; font-family: 'Courier New', monospace; font-size: 13px; resize: vertical; min-height: 80px; line-height: 1.6; }
    .sql-run { background: #7dd3fc20; border: 1px solid #7dd3fc50; color: #7dd3fc; border-radius: 8px; padding: 8px 20px; font-size: 13px; font-weight: 700; cursor: pointer; font-family: inherit; transition: all .2s; }
    .sql-run:hover { background: #7dd3fc30; }
    .sql-result { font-size: 12px; color: #a3e635; margin-top: 10px; font-family: monospace; line-height: 1.6; white-space: pre-wrap; word-break: break-all; }

    /* ── PAGES ── */
    .page { display: none; }
    .page.active { display: block; }

    /* ── HERO ── */
    .hero { max-width: 1200px; margin: 0 auto; padding: 60px 24px; display: grid; grid-template-columns: 1fr 1fr; gap: 60px; align-items: center; }
    .hero-badge { display: inline-flex; align-items: center; gap: 6px; background: var(--green-light); color: var(--green-dark); font-size: 12px; font-weight: 700; padding: 6px 14px; border-radius: 999px; margin-bottom: 20px; }
    .hero h1 { font-size: 42px; font-weight: 900; line-height: 1.2; letter-spacing: -1.5px; margin-bottom: 16px; }
    .hero h1 em { color: var(--green); font-style: normal; }
    .hero p.sub { color: var(--muted); font-size: 15px; line-height: 1.7; margin-bottom: 28px; }
    .input-row { display: flex; gap: 10px; margin-bottom: 12px; }
    .input-row input { flex: 1; border: 2px solid #e5e7eb; border-radius: 14px; padding: 14px 18px; font-size: 15px; font-family: inherit; outline: none; background: #fafafa; transition: border .2s; }
    .input-row input:focus { border-color: var(--green); background: #fff; }
    .btn-add { background: #1a1a2e; color: #fff; border: none; border-radius: 14px; padding: 14px 22px; font-size: 14px; font-weight: 700; cursor: pointer; white-space: nowrap; font-family: inherit; transition: background .2s; }
    .btn-add:hover { background: #2d2d4e; }
    .basket-box { background: #f8faf9; border: 1px solid #e5e7eb; border-radius: 14px; padding: 14px 18px; margin-bottom: 20px; }
    .basket-label { font-size: 11px; font-weight: 700; color: #aaa; letter-spacing: .5px; text-transform: uppercase; margin-bottom: 8px; }
    .chips { display: flex; flex-wrap: wrap; gap: 8px; }
    .chip { background: var(--green-light); color: var(--green-dark); font-size: 12px; font-weight: 600; padding: 5px 12px; border-radius: 999px; border: 1px solid #b8f0d4; display: flex; align-items: center; gap: 6px; }
    .chip-remove { cursor: pointer; font-size: 14px; line-height: 1; opacity: .6; }
    .chip-remove:hover { opacity: 1; }
    .btn-main { width: 100%; background: var(--green); color: #fff; border: none; border-radius: 14px; padding: 16px; font-size: 16px; font-weight: 700; cursor: pointer; font-family: inherit; transition: background .2s, transform .1s; }
    .btn-main:hover { background: var(--green-dark); transform: translateY(-1px); }
    .hero-img { border-radius: 28px; overflow: hidden; aspect-ratio: 1; box-shadow: 0 30px 60px rgba(0,0,0,.12); }
    .hero-img img { width: 100%; height: 100%; object-fit: cover; }

    /* ── POPULAR ── */
    .popular { max-width: 1200px; margin: 0 auto 60px; padding: 0 24px; }
    .section-title { font-size: 22px; font-weight: 800; margin-bottom: 20px; letter-spacing: -.5px; display: flex; justify-content: space-between; align-items: center; }
    .cards-row { display: grid; grid-template-columns: repeat(3, 1fr); gap: 20px; }
    .recipe-card { border-radius: 20px; overflow: hidden; background: #fff; box-shadow: 0 4px 16px rgba(0,0,0,.06); transition: transform .2s, box-shadow .2s; cursor: pointer; }
    .recipe-card:hover { transform: translateY(-4px); box-shadow: 0 12px 32px rgba(0,0,0,.12); }
    .recipe-card img { width: 100%; height: 180px; object-fit: cover; }
    .recipe-card-body { padding: 16px 18px; }
    .recipe-card-body h3 { font-size: 16px; font-weight: 700; margin-bottom: 4px; }
    .recipe-card-body p { font-size: 13px; color: var(--muted); }
    .calorie-info { font-size: 12px; color: var(--accent); font-weight: 700; margin-top: 6px; }

    /* ── ROULETTE ── */
    .roulette-page { max-width: 1200px; margin: 0 auto; padding: 40px 24px; }
    .roulette-header { text-align: center; margin-bottom: 36px; }
    .roulette-header h2 { font-size: 30px; font-weight: 900; letter-spacing: -1px; margin-bottom: 8px; }
    .roulette-header p { color: var(--muted); font-size: 14px; }
    .roulette-layout { display: grid; grid-template-columns: 1fr 380px; gap: 40px; align-items: start; }
    .roulette-center { display: flex; flex-direction: column; align-items: center; gap: 24px; }
    .canvas-wrap { position: relative; width: 340px; height: 340px; }
    #rouletteCanvas { border-radius: 50%; box-shadow: 0 16px 48px rgba(43,182,115,.25); display: block; }
    .roulette-arrow { position: absolute; top: -18px; left: 50%; transform: translateX(-50%); width: 0; height: 0; border-left: 14px solid transparent; border-right: 14px solid transparent; border-top: 26px solid var(--accent); filter: drop-shadow(0 3px 6px rgba(255,107,53,.4)); z-index: 10; }
    .roulette-btns { display: flex; gap: 12px; }
    .btn-spin { background: var(--green); color: #fff; border: none; border-radius: 14px; padding: 14px 32px; font-size: 16px; font-weight: 800; cursor: pointer; font-family: inherit; transition: background .2s, transform .1s; box-shadow: 0 6px 20px rgba(43,182,115,.3); }
    .btn-spin:hover:not(:disabled) { background: var(--green-dark); transform: translateY(-2px); }
    .btn-spin:disabled { opacity: .5; cursor: not-allowed; }
    .btn-reset { background: #f1f3f2; color: #555; border: none; border-radius: 14px; padding: 14px 24px; font-size: 15px; font-weight: 600; cursor: pointer; font-family: inherit; }
    .btn-reset:hover { background: #e5e7e6; }
    .recipe-side { background: #fff; border-radius: 24px; padding: 24px; box-shadow: 0 4px 20px rgba(0,0,0,.06); }
    .recipe-side h3 { font-size: 16px; font-weight: 800; margin-bottom: 16px; }
    .side-list { display: flex; flex-direction: column; gap: 10px; max-height: 500px; overflow-y: auto; padding-right: 4px; }
    .side-card { background: #f8faf9; border: 2px solid #eef2f1; border-radius: 14px; padding: 14px 16px; cursor: pointer; transition: all .2s; }
    .side-card:hover { border-color: var(--green); background: #f1fff7; }
    .side-card.active { border-color: var(--green); background: var(--green); }
    .side-card.active .side-name, .side-card.active .side-parts { color: #fff !important; }
    .side-name { font-size: 15px; font-weight: 700; margin-bottom: 4px; }
    .side-parts { font-size: 12px; color: #999; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }

    /* ── RECIPE MODAL ── */
    .modal-overlay { display: none; position: fixed; inset: 0; background: rgba(0,0,0,.55); z-index: 1000; align-items: center; justify-content: center; padding: 20px; backdrop-filter: blur(4px); }
    .modal-overlay.open { display: flex; }
    .modal { background: #fff; border-radius: 28px; max-width: 800px; width: 100%; max-height: 90vh; overflow-y: auto; box-shadow: 0 40px 80px rgba(0,0,0,.2); }
    .modal-top { padding: 32px 36px 24px; border-bottom: 1px solid #f0f0f0; display: flex; justify-content: space-between; align-items: flex-start; }
    .modal-title { font-size: 28px; font-weight: 900; letter-spacing: -1px; }
    .modal-close { background: #f1f3f2; border: none; border-radius: 50%; width: 36px; height: 36px; font-size: 18px; cursor: pointer; display: flex; align-items: center; justify-content: center; }
    .modal-close:hover { background: #e5e7e6; }
    .modal-body { padding: 24px 36px 36px; }
    .calorie-tag { display: inline-block; background: #fff0f0; color: #ff5c5c; font-size: 13px; font-weight: 700; padding: 5px 14px; border-radius: 999px; margin-bottom: 16px; }
    .ingredients-block { background: #f8faf9; border-radius: 14px; padding: 16px 20px; font-size: 14px; line-height: 1.7; color: #444; margin-bottom: 24px; }
    .modal-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 24px; }
    .modal-img { border-radius: 18px; overflow: hidden; aspect-ratio: 1; }
    .modal-img img { width: 100%; height: 100%; object-fit: cover; }
    .steps-list { display: flex; flex-direction: column; gap: 10px; overflow-y: auto; max-height: 300px; }
    .step { display: flex; gap: 10px; font-size: 14px; line-height: 1.6; color: #333; }
    .step-num { color: var(--green); font-weight: 800; flex-shrink: 0; }

    /* ── CALENDAR ── */
    .calendar-page { max-width: 1200px; margin: 0 auto; padding: 40px 24px; }
    .cal-layout { display: grid; grid-template-columns: 1fr 340px; gap: 32px; align-items: start; }
    .cal-box { background: #fff; border-radius: 24px; padding: 28px; box-shadow: 0 4px 20px rgba(0,0,0,.06); }
    .cal-nav { display: flex; align-items: center; justify-content: space-between; margin-bottom: 20px; }
    .cal-nav-btn { background: #f1f3f2; border: none; border-radius: 10px; width: 36px; height: 36px; font-size: 16px; cursor: pointer; display: flex; align-items: center; justify-content: center; }
    .cal-nav-btn:hover { background: var(--green-light); }
    .cal-month { font-size: 18px; font-weight: 800; }
    .cal-grid { display: grid; grid-template-columns: repeat(7,1fr); gap: 4px; }
    .cal-day-label { text-align: center; font-size: 11px; font-weight: 700; color: #aaa; padding: 6px 0; }
    .cal-cell { aspect-ratio: 1; border-radius: 10px; display: flex; flex-direction: column; align-items: center; justify-content: center; cursor: pointer; font-size: 13px; font-weight: 500; transition: background .15s; position: relative; gap: 2px; }
    .cal-cell:hover { background: var(--green-light); }
    .cal-cell.today { background: var(--green); color: #fff; font-weight: 800; }
    .cal-cell.other-month { color: #ccc; }
    .dot-row { display: flex; gap: 2px; justify-content: center; flex-wrap: wrap; max-width: 30px; }
    .expiry-dot { width: 6px; height: 6px; border-radius: 50%; flex-shrink: 0; }
    .side-panel { background: #fff; border-radius: 24px; padding: 24px; box-shadow: 0 4px 20px rgba(0,0,0,.06); }
    .side-panel h3 { font-size: 16px; font-weight: 800; margin-bottom: 16px; }
    .add-item-form { display: flex; flex-direction: column; gap: 10px; margin-bottom: 20px; }
    .add-item-form input { border: 2px solid #e5e7eb; border-radius: 12px; padding: 10px 14px; font-size: 14px; font-family: inherit; outline: none; transition: border .2s; }
    .add-item-form input:focus { border-color: var(--green); }
    .btn-sm { background: var(--green); color: #fff; border: none; border-radius: 12px; padding: 10px 16px; font-size: 14px; font-weight: 700; cursor: pointer; font-family: inherit; }
    .btn-sm:hover { background: var(--green-dark); }
    .items-list { display: flex; flex-direction: column; gap: 8px; max-height: 400px; overflow-y: auto; }
    .expiry-item { background: #f8faf9; border-radius: 12px; padding: 12px 14px; display: flex; align-items: center; gap: 10px; border-left: 4px solid; }
    .expiry-item.safe { border-color: var(--green); }
    .expiry-item.warn { border-color: #f59e0b; background: #fffbeb; }
    .expiry-item.danger { border-color: #ef4444; background: #fef2f2; }
    .expiry-item.expired { border-color: #9ca3af; background: #f9fafb; opacity: .7; }
    .item-name { font-size: 14px; font-weight: 700; flex: 1; }
    .item-date { font-size: 12px; color: var(--muted); }
    .item-badge { font-size: 11px; font-weight: 700; padding: 3px 8px; border-radius: 999px; }
    .safe .item-badge { background: var(--green-light); color: var(--green-dark); }
    .warn .item-badge { background: #fef3c7; color: #92400e; }
    .danger .item-badge { background: #fee2e2; color: #991b1b; }
    .expired .item-badge { background: #f3f4f6; color: #6b7280; }
    .item-del { background: none; border: none; color: #ccc; cursor: pointer; font-size: 16px; padding: 2px; }
    .item-del:hover { color: #ef4444; }
    .alert-banner { background: linear-gradient(135deg, #ff6b35, #ff8c5a); color: #fff; border-radius: 16px; padding: 16px 20px; margin-bottom: 24px; display: flex; align-items: center; gap: 12px; }
    .alert-banner.hidden { display: none; }

    /* ── SHOPPING PAGE ── */
    .shopping-page { max-width: 1100px; margin: 0 auto; padding: 60px 24px; }
    .shopping-page h2 { font-size: 30px; font-weight: 900; letter-spacing: -1px; margin-bottom: 8px; }
    .shopping-page .sub-desc { color: var(--muted); font-size: 15px; margin-bottom: 48px; }

    /* 3열 그리드로 변경 */
    .shop-cards { display: grid; grid-template-columns: repeat(3, 1fr); gap: 20px; }

    .shop-card { border-radius: 24px; padding: 30px 28px; text-decoration: none; color: var(--text); display: flex; flex-direction: column; align-items: flex-start; gap: 12px; box-shadow: 0 6px 24px rgba(0,0,0,.08); transition: transform .2s, box-shadow .2s; cursor: pointer; border: 2px solid transparent; }
    .shop-card:hover { transform: translateY(-5px); box-shadow: 0 16px 40px rgba(0,0,0,.14); border-color: currentColor; }

    /* 기존 2개 */
    .shop-card.coupang { background: linear-gradient(135deg, #fff5f0 0%, #ffe8dc 100%); }
    .shop-card.kurly   { background: linear-gradient(135deg, #f5f0ff 0%, #e8dcff 100%); }

    /* 신규 6개 */
    .shop-card.ssg      { background: linear-gradient(135deg, #fff8f0 0%, #ffecd8 100%); }
    .shop-card.lotteon  { background: linear-gradient(135deg, #fff0f0 0%, #ffd6d6 100%); }
    .shop-card.nonghyup { background: linear-gradient(135deg, #f0fff4 0%, #d6ffe6 100%); }
    .shop-card.eleventh { background: linear-gradient(135deg, #fff0f5 0%, #ffd6e8 100%); }
    .shop-card.gmarket  { background: linear-gradient(135deg, #fffff0 0%, #fff6c8 100%); }
    .shop-card.oasis    { background: linear-gradient(135deg, #f0faff 0%, #d6f0ff 100%); }

    .shop-card-icon { font-size: 40px; line-height: 1; }
    .shop-card-name { font-size: 22px; font-weight: 900; letter-spacing: -.5px; }

    .shop-card.coupang  .shop-card-name { color: #c0392b; }
    .shop-card.kurly    .shop-card-name { color: #6c2bd9; }
    .shop-card.ssg      .shop-card-name { color: #c4551a; }
    .shop-card.lotteon  .shop-card-name { color: #c0181b; }
    .shop-card.nonghyup .shop-card-name { color: #1a7a3e; }
    .shop-card.eleventh .shop-card-name { color: #d4145a; }
    .shop-card.gmarket  .shop-card-name { color: #b58800; }
    .shop-card.oasis    .shop-card-name { color: #0074b5; }

    .shop-card-desc { font-size: 13px; color: var(--muted); line-height: 1.6; }
    .shop-card-btn { margin-top: 6px; padding: 10px 22px; border-radius: 999px; font-size: 13px; font-weight: 800; border: none; cursor: pointer; font-family: inherit; transition: filter .2s; }
    .shop-card-btn:hover { filter: brightness(1.12); }

    .shop-card.coupang  .shop-card-btn { background: #c0392b; color: #fff; }
    .shop-card.kurly    .shop-card-btn { background: #6c2bd9; color: #fff; }
    .shop-card.ssg      .shop-card-btn { background: #c4551a; color: #fff; }
    .shop-card.lotteon  .shop-card-btn { background: #c0181b; color: #fff; }
    .shop-card.nonghyup .shop-card-btn { background: #1a7a3e; color: #fff; }
    .shop-card.eleventh .shop-card-btn { background: #d4145a; color: #fff; }
    .shop-card.gmarket  .shop-card-btn { background: #b58800; color: #fff; }
    .shop-card.oasis    .shop-card-btn { background: #0074b5; color: #fff; }

    @media (max-width: 900px) { .shop-cards { grid-template-columns: repeat(2, 1fr); } }
    @media (max-width: 600px) { .shop-cards { grid-template-columns: 1fr; } }

    /* ── RECIPE LIST TABLE ── */
    .recipe-list-table { display: flex; flex-direction: column; gap: 10px; }
    .recipe-row { background: #fff; border-radius: 16px; padding: 18px 24px; display: flex; align-items: center; gap: 20px; box-shadow: 0 2px 10px rgba(0,0,0,.05); cursor: pointer; transition: transform .15s, box-shadow .15s, border-color .15s; border: 2px solid transparent; }
    .recipe-row:hover { transform: translateX(4px); box-shadow: 0 6px 20px rgba(43,182,115,.15); border-color: var(--green); }
    .recipe-row-num { width: 32px; height: 32px; border-radius: 50%; background: var(--green-light); color: var(--green-dark); font-size: 13px; font-weight: 800; display: flex; align-items: center; justify-content: center; flex-shrink: 0; }
    .recipe-row-info { flex: 1; }
    .recipe-row-name { font-size: 16px; font-weight: 700; margin-bottom: 4px; }
    .recipe-row-parts { font-size: 13px; color: var(--muted); }
    .recipe-row-kcal { font-size: 13px; font-weight: 700; color: var(--accent); white-space: nowrap; }
    .recipe-row-arrow { color: #ccc; font-size: 18px; }
    .no-recipe-result { text-align: center; color: var(--muted); padding: 48px 0; font-size: 15px; }

    /* ── POPULAR (홈 카드) ── */
    .recipe-card-noimg { border-radius: 20px; overflow: hidden; background: #fff; box-shadow: 0 4px 16px rgba(0,0,0,.06); transition: transform .2s, box-shadow .2s; cursor: pointer; border: 2px solid transparent; }
    .recipe-card-noimg:hover { transform: translateY(-4px); box-shadow: 0 12px 32px rgba(0,0,0,.12); border-color: var(--green); }
    .recipe-card-noimg-body { padding: 28px 22px; }
    .recipe-card-noimg-icon { font-size: 36px; margin-bottom: 10px; }
    .recipe-card-noimg-body h3 { font-size: 17px; font-weight: 800; margin-bottom: 6px; }
    .recipe-card-noimg-body p { font-size: 13px; color: var(--muted); }

    /* ── LOADING ── */
    .loading-overlay { display: none; position: fixed; inset: 0; background: rgba(255,255,255,.85); z-index: 500; align-items: center; justify-content: center; flex-direction: column; gap: 16px; backdrop-filter: blur(4px); }
    .loading-overlay.show { display: flex; }
    .spinner { width: 48px; height: 48px; border: 5px solid var(--green-light); border-top-color: var(--green); border-radius: 50%; animation: spin .8s linear infinite; }
    @keyframes spin { to { transform: rotate(360deg); } }
    .loading-text { color: var(--green-dark); font-weight: 700; font-size: 15px; }

    /* ── MISC ── */
    @keyframes fadeIn { from { opacity: 0; transform: translateY(12px); } to { opacity: 1; transform: translateY(0); } }
    .fade-in { animation: fadeIn .3s ease; }
    @media (max-width: 900px) { .hero, .roulette-layout, .cal-layout, .modal-grid { grid-template-columns: 1fr; } .hero-img { display: none; } .cards-row { grid-template-columns: 1fr 1fr; } }
    @media (max-width: 600px) { .hero h1 { font-size: 30px; } .cards-row { grid-template-columns: 1fr; } nav { display: none; } }
  </style>
</head>
<body>

<header>
  <div class="header-inner">
    <div class="logo" onclick="showPage('home', document.getElementById('menu-home'))">냉장<span>Go</span></div>
    <nav>
      <a href="#" class="nav-menu active" id="menu-home" onclick="showPage('home', this)">홈</a>
      <a href="#" class="nav-menu" id="menu-recipe" onclick="showPage('recipe-list-page', this)">레시피</a>
      <a href="#" class="nav-menu" id="menu-calendar" onclick="showPage('calendar', this)">냉장고 관리</a>
      <a href="#" class="nav-menu" id="menu-shopping" onclick="showPage('shopping', this)">🛒 장보기</a>
    </nav>
    <div class="header-auth" id="headerAuth">
      <button class="login-btn outline" onclick="openAuth('login')">로그인</button>
      <button class="login-btn" onclick="openAuth('signup')">회원가입</button>
    </div>
  </div>
</header>

<!-- ══ AUTH MODAL ══ -->
<div class="auth-overlay" id="authOverlay" onclick="closeAuthIfOverlay(event)">
  <div class="auth-box fade-in">
    <button class="auth-close" onclick="closeAuth()">✕</button>
    <div class="auth-logo">냉장<span style="color:#1a1a2e">Go</span></div>
    <div class="auth-sub" id="authSub">냉장고 관리를 시작해 보세요</div>
    <div class="auth-tabs">
      <button class="auth-tab active" id="tabLogin" onclick="switchTab('login')">로그인</button>
      <button class="auth-tab" id="tabSignup" onclick="switchTab('signup')">회원가입</button>
    </div>
    <div class="auth-error" id="authError"></div>
    <div id="formLogin">
      <div class="auth-form">
        <label>이메일</label>
        <input type="email" id="loginEmail" placeholder="example@email.com">
        <label>비밀번호</label>
        <input type="password" id="loginPw" placeholder="비밀번호 입력" onkeyup="if(event.key==='Enter')doLogin()">
        <button class="auth-submit" onclick="doLogin()">로그인</button>
      </div>
    </div>
    <div id="formSignup" style="display:none">
      <div class="auth-form">
        <label>닉네임</label>
        <input type="text" id="signupNick" placeholder="냉장고마스터" maxlength="12">
        <label>이메일</label>
        <input type="email" id="signupEmail" placeholder="example@email.com">
        <label>비밀번호 <span style="color:#aaa;font-weight:400">(6자 이상)</span></label>
        <input type="password" id="signupPw" placeholder="••••••••" onkeyup="if(event.key==='Enter')doSignup()">
        <label>비밀번호 확인</label>
        <input type="password" id="signupPw2" placeholder="••••••••" onkeyup="if(event.key==='Enter')doSignup()">
        <button class="auth-submit" onclick="doSignup()">가입하기</button>
      </div>
    </div>
  </div>
</div>

<!-- ══ DB VIEWER PANEL ══ -->
<div class="db-overlay" id="dbOverlay" onclick="closeDbIfOverlay(event)">
  <div class="db-panel fade-in">
    <div class="db-header">
      <div class="db-title">🗄️ 냉장Go — IndexedDB 관리자 패널</div>
      <div style="display:flex;align-items:center;gap:12px">
        <div class="db-tabs">
          <button class="db-tab active" onclick="switchDbTab('users')">users</button>
          <button class="db-tab" onclick="switchDbTab('ingredients')">ingredients</button>
          <button class="db-tab" onclick="switchDbTab('expiry')">expiry_items</button>
          <button class="db-tab" onclick="switchDbTab('sql')">SQL 쿼리</button>
        </div>
        <button class="db-close" onclick="closeDbPanel()">닫기</button>
      </div>
    </div>
    <div class="db-body" id="dbBody">
      <div class="db-empty">로딩 중...</div>
    </div>
  </div>
</div>

<!-- ══ RECIPE MODAL ══ -->
<div class="modal-overlay" id="recipeModal" onclick="closeModalIfOverlay(event)">
  <div class="modal fade-in">
    <div class="modal-top">
      <div>
        <div class="modal-title" id="modalTitle">레시피 이름</div>
        <div class="calorie-tag" id="modalCalorie">🔥 0 kcal</div>
      </div>
      <button class="modal-close" onclick="closeModal()">✕</button>
    </div>
    <div class="modal-body">
      <div class="ingredients-block" id="modalIngredients">재료 정보</div>
      <div id="modalSteps" style="display:flex;flex-direction:column;gap:10px;"></div>
    </div>
  </div>
</div>

<!-- ══ LOADING ══ -->
<div class="loading-overlay" id="loadingOverlay">
  <div class="spinner"></div>
  <div class="loading-text">레시피 검색 중...</div>
</div>

<!-- ══════════ PAGES ══════════ -->
<div class="page active" id="page-home">

  <section class="hero">

    <div>
      <div class="hero-badge">🥬 냉장고 속 재료로 오늘 뭐 먹지?</div>

      <h1>
        <em>재료</em>는 있는데,<br>
        메뉴가 안 떠오를 때
      </h1>

      <p class="sub">
        냉장고 속 재료를 입력하면<br>
        지금 만들 수 있는 요리를 추천해드립니다.
      </p>

      <div class="input-row">

        <input
          type="text"
          id="ingredientInput"
          placeholder="예: 김치, 달걀, 두부"
          onkeyup="if(event.key==='Enter') addIngredient()"
        >

        <button class="btn-add" onclick="addIngredient()">
          재료 등록하기
        </button>

      </div>

      <div class="basket-box">
        <div class="basket-label">내 냉장고 현황</div>

        <div class="chips" id="chipContainer">
          <span style="color:#aaa;font-size:13px;">
            등록된 재료가 없습니다.
          </span>
        </div>
      </div>

      <button class="btn-main" onclick="fetchRecipesAndGo()">
        🚀 레시피 추천받기
      </button>
    </div>

    <!-- 냉장고 영역 -->
    <div class="hero-img">

      <div class="fridge-container">

        <!-- 냉장고 이미지 -->
        <img src="냉장고.png" class="fridge-img" alt="냉장고">

        <!-- 재료 표시 영역 -->
        <div id="ingredient-layer"></div>

      </div>

    </div>

  </section>

  <div class="popular">
    <div class="section-title">🔥 실시간 추천 인기 레시피</div>

    <div class="cards-row" id="popularCardsContainer"></div>
  </div>

</div>

<style>

/* 냉장고 영역 */
.fridge-container{
  position:relative;
  width:420px;
  margin:auto;
}

/* 냉장고 이미지 */
.fridge-img{
  width:100%;
  display:block;
  border-radius:20px;
}

/* 재료 표시 영역 */
#ingredient-layer{
  position:absolute;

  top:80px;
  left:145px;

  width:180px;
  height:500px;

  display:flex;
  flex-wrap:wrap;
  align-content:flex-start;

  gap:10px;
  padding:10px;
}

/* 재료 아이템 */
.ingredient-item{
  background:rgba(255,255,255,0.9);

  border-radius:12px;

  padding:6px 10px;

  font-size:14px;

  box-shadow:0 2px 5px rgba(0,0,0,0.15);

  animation:pop 0.3s ease;
}

/* 애니메이션 */
@keyframes pop{

  from{
    transform:scale(0.5);
    opacity:0;
  }

  to{
    transform:scale(1);
    opacity:1;
  }
}

</style>

<script>
  
/* 냉장고 재료 출력 */
function renderIngredients(){

  const layer =
    document.getElementById("ingredient-layer");

  layer.innerHTML = "";

  ingredients.forEach(item => {

    const div = document.createElement("div");

    div.className = "ingredient-item";

    div.innerText = item;

    layer.appendChild(div);

  });
}
</script>

  <div class="popular">
    <div class="section-title">🔥 실시간 추천 인기 레시피</div>
    <div class="cards-row" id="popularCardsContainer"></div>
  </div>
</div>

<div class="page" id="page-recipe-list-page">
  <div class="popular">
    <div class="section-title">
      <span>🍳 오늘 뭐 먹지?</span>
      <span style="font-size:14px;color:var(--muted);font-weight:normal;">칼로리 낮은 순</span>
    </div>
    <div style="margin-bottom:24px;">
      <div style="position:relative;max-width:480px;">
        <input type="text" id="recipeSearchInput" placeholder="🔍 레시피명 또는 재료로 검색..." oninput="filterRecipes()" style="width:100%;border:2px solid #e5e7eb;border-radius:14px;padding:13px 18px 13px 46px;font-size:15px;font-family:inherit;outline:none;background:#fff;transition:border .2s;">
        <span style="position:absolute;left:16px;top:50%;transform:translateY(-50%);font-size:17px;pointer-events:none;">🔍</span>
      </div>
      <div id="recipeSearchCount" style="font-size:13px;color:var(--muted);margin-top:8px;"></div>
    </div>
    <div class="recipe-list-table" id="allRecipesCardsContainer"></div>
  </div>
</div>

<div class="page" id="page-roulette">
  <div class="roulette-page fade-in">
    <div class="roulette-header">
      <h2>🎯 오늘 저녁 추천 메뉴 룰렛!</h2>
      <p>룰렛을 돌려 오늘의 메뉴를 정해보세요!</p>
    </div>
    <div class="roulette-layout">
      <div class="roulette-center">
        <div class="canvas-wrap">
          <div class="roulette-arrow"></div>
          <canvas id="rouletteCanvas" width="340" height="340"></canvas>
        </div>
        <div class="roulette-btns">
          <button class="btn-spin" id="spinBtn" onclick="spinRoulette()">🔮 룰렛 돌리기!!</button>
          <button class="btn-reset" onclick="resetAll()">← 다시 입력</button>
        </div>
        <p style="font-size:13px;color:#aaa;">룰렛이 멈추면 레시피가 자동으로 표시됩니다</p>
      </div>
      <div class="recipe-side">
        <h3>📋 추천 레시피 (<span id="recipeCount">0</span>개)</h3>
        <div class="side-list" id="sideList"><div style="color:#aaa;font-size:13px;text-align:center;padding:20px 0;">로딩 중...</div></div>
      </div>
    </div>
  </div>
</div>

<div class="page" id="page-calendar">
  <div class="calendar-page fade-in">
    <div style="display:flex;align-items:center;justify-content:space-between;margin-bottom:28px;">
      <h2 style="font-size:26px;font-weight:900;letter-spacing:-.5px">🗓️ 냉장고 유통기한 관리</h2>
    </div>
    <div class="alert-banner hidden" id="alertBanner">
      <div style="font-size:22px">⚠️</div>
      <div style="flex:1">
        <strong id="alertTitle" style="display:block;font-weight:800;margin-bottom:2px;font-size:15px"></strong>
        <span id="alertBody" style="font-size:14px"></span>
      </div>
    </div>
    <div class="cal-layout">
      <div class="cal-box">
        <div class="cal-nav">
          <button class="cal-nav-btn" onclick="changeMonth(-1)">◀</button>
          <span class="cal-month" id="calMonthLabel"></span>
          <button class="cal-nav-btn" onclick="changeMonth(1)">▶</button>
        </div>
        <div class="cal-grid" id="calGrid"></div>
      </div>
      <div class="side-panel">
        <h3>➕ 재료 등록</h3>
        <div class="add-item-form">
          <input type="text" id="itemName" placeholder="재료 이름 (예: 우유)">
          <input type="date" id="itemExpiry">
          <button class="btn-sm" onclick="addExpiryItem()">등록하기</button>
        </div>
        <h3>📦 등록된 재료</h3>
        <div class="items-list" id="itemsList">
          <div style="color:#aaa;font-size:13px;text-align:center;padding:16px 0;">등록된 재료가 없습니다.</div>
        </div>
      </div>
    </div>
  </div>
</div>

<!-- ══ 장보기 페이지 — 8개 쇼핑몰 ══ -->
<div class="page" id="page-shopping">
  <div class="shopping-page fade-in">
    <h2>🛒 장보기</h2>
    <p class="sub-desc">원하는 쇼핑몰을 선택하면 바로 이동합니다.</p>
    <div class="shop-cards">

      <!-- 쿠팡 -->
      <a href="https://www.coupang.com/?src=1042001&spec=10304102&addtag=900&ctag=HOME&lptag=%EC%BF%A0%ED%8C%A1%EC%82%AC%EC%9D%B4%ED%8A%B8&itime=20260602101923&pageType=HOME&pageValue=HOME&wPcid=17803631636520658783979&wRef=www.bing.com&wTime=20260602101923&redirect=landing&mcid=4a1d85e882794c658255cce6ea6f862f&n_match=1&n_keyword=%EC%BF%A0%ED%8C%A1%EC%82%AC%EC%9D%B4%ED%8A%B8&n_network=&n_ad_group=grp-a001-01-000000009307622&n_ad=nad-a001-01-000000255553660&n_rank=1&n_keyword_id=nkw-a001-01-000005716981265&n_media=335738&n_campaign_type=1&n_query=%EC%BF%A0%ED%8C%A1%EC%82%AC%EC%9D%B4%ED%8A%B8" target="_blank" class="shop-card coupang">
        <div class="shop-card-icon">🛍️</div>
        <div class="shop-card-name">쿠팡</div>
        <div class="shop-card-desc">로켓배송으로 신선식품·생필품을<br>빠르게 받아보세요.</div>
        <button class="shop-card-btn">바로가기 →</button>
      </a>

      <!-- 마켓컬리 -->
      <a href="https://www.kurly.com/main" target="_blank" class="shop-card kurly">
        <div class="shop-card-icon">🌿</div>
        <div class="shop-card-name">마켓컬리</div>
        <div class="shop-card-desc">신선한 유기농·프리미엄 식재료를<br>새벽배송으로 만나보세요.</div>
        <button class="shop-card-btn">바로가기 →</button>
      </a>

      <!-- SSG.COM -->
      <a href="https://www.ssg.com/" target="_blank" class="shop-card ssg">
        <div class="shop-card-icon">🏬</div>
        <div class="shop-card-name">SSG.COM</div>
        <div class="shop-card-desc">신세계·이마트 상품을 한 곳에서,<br>쓱배송으로 빠르게 받아보세요.</div>
        <button class="shop-card-btn">바로가기 →</button>
      </a>

      <!-- 롯데ON -->
      <a href="https://www.lotteon.com/p/display/main/lotteon" target="_blank" class="shop-card lotteon">
        <div class="shop-card-icon">🎰</div>
        <div class="shop-card-name">롯데ON</div>
        <div class="shop-card-desc">롯데 계열 브랜드 상품과<br>다양한 식품을 한번에 쇼핑하세요.</div>
        <button class="shop-card-btn">바로가기 →</button>
      </a>

      <!-- 농협몰 -->
      <a href="https://www.nonghyupmall.com/BC31010R/main.nh?emdvEndYn=Y&basketCnt=0&cdnAplYn=N&nhVuchYn=N" target="_blank" class="shop-card nonghyup">
        <div class="shop-card-icon">🌾</div>
        <div class="shop-card-name">농협몰</div>
        <div class="shop-card-desc">농협 직거래 신선 농산물·가공식품을<br>산지 직송으로 받아보세요.</div>
        <button class="shop-card-btn">바로가기 →</button>
      </a>

      <!-- 11번가 -->
      <a href="https://www.11st.co.kr/" target="_blank" class="shop-card eleventh">
        <div class="shop-card-icon">1️⃣</div>
        <div class="shop-card-name">11번가</div>
        <div class="shop-card-desc">다양한 브랜드 식품과 생활용품을<br>합리적인 가격으로 구매하세요.</div>
        <button class="shop-card-btn">바로가기 →</button>
      </a>

      <!-- G마켓 -->
      <a href="https://www.gmarket.co.kr/" target="_blank" class="shop-card gmarket">
        <div class="shop-card-icon">🅶</div>
        <div class="shop-card-name">G마켓</div>
        <div class="shop-card-desc">Super Delivery로 신선식품부터<br>가공식품까지 빠르게 배송됩니다.</div>
        <button class="shop-card-btn">바로가기 →</button>
      </a>

      <!-- 오아시스 -->
      <a href="https://www.oasis.co.kr/main#" target="_blank" class="shop-card oasis">
        <div class="shop-card-icon">🌴</div>
        <div class="shop-card-name">오아시스</div>
        <div class="shop-card-desc">친환경·유기농 신선식품 전문몰,<br>건강한 식탁을 위한 선택.</div>
        <button class="shop-card-btn">바로가기 →</button>
      </a>

    </div>
  </div>
</div>

<script>
const DB_NAME = 'NaengJangGoDB';
const DB_VERSION = 3;
let db = null;

function openDB() {
  return new Promise((resolve, reject) => {
    if (db) { resolve(db); return; }
    const req = indexedDB.open(DB_NAME, DB_VERSION);
    req.onupgradeneeded = e => {
      const d = e.target.result;
      if (!d.objectStoreNames.contains('users')) {
        const us = d.createObjectStore('users', { keyPath: 'id', autoIncrement: true });
        us.createIndex('email', 'email', { unique: true });
      }
      if (!d.objectStoreNames.contains('ingredients')) {
        const is = d.createObjectStore('ingredients', { keyPath: 'id', autoIncrement: true });
        is.createIndex('userId', 'userId', { unique: false });
      }
      if (!d.objectStoreNames.contains('expiry_items')) {
        const es = d.createObjectStore('expiry_items', { keyPath: 'id', autoIncrement: true });
        es.createIndex('userId', 'userId', { unique: false });
      }
    };
    req.onsuccess = e => { db = e.target.result; resolve(db); };
    req.onerror = e => reject(e.target.error);
  });
}

function dbAdd(store, data) {
  return new Promise((resolve, reject) => {
    const tx = db.transaction(store, 'readwrite');
    const req = tx.objectStore(store).add(data);
    req.onsuccess = e => resolve(e.target.result);
    req.onerror = e => reject(e.target.error);
  });
}
function dbGetAll(store) {
  return new Promise((resolve, reject) => {
    const tx = db.transaction(store, 'readonly');
    const req = tx.objectStore(store).getAll();
    req.onsuccess = e => resolve(e.target.result);
    req.onerror = e => reject(e.target.error);
  });
}
function dbGetByIndex(store, index, value) {
  return new Promise((resolve, reject) => {
    const tx = db.transaction(store, 'readonly');
    const req = tx.objectStore(store).index(index).getAll(value);
    req.onsuccess = e => resolve(e.target.result);
    req.onerror = e => reject(e.target.error);
  });
}
function dbDelete(store, id) {
  return new Promise((resolve, reject) => {
    const tx = db.transaction(store, 'readwrite');
    const req = tx.objectStore(store).delete(id);
    req.onsuccess = () => resolve();
    req.onerror = e => reject(e.target.error);
  });
}
function dbGetByEmail(email) {
  return new Promise((resolve, reject) => {
    const tx = db.transaction('users', 'readonly');
    const req = tx.objectStore('users').index('email').get(email);
    req.onsuccess = e => resolve(e.target.result);
    req.onerror = e => reject(e.target.error);
  });
}

function simpleHash(str) {
  let h = 5381;
  for (let i = 0; i < str.length; i++) h = ((h << 5) + h) ^ str.charCodeAt(i);
  return 'h_' + (h >>> 0).toString(16);
}

let currentUser = JSON.parse(sessionStorage.getItem('naengUser') || 'null');

function setUser(user) {
  currentUser = user;
  sessionStorage.setItem('naengUser', JSON.stringify(user));
  renderHeaderAuth();
  loadUserData();
}

function renderHeaderAuth() {
  const el = document.getElementById('headerAuth');
  if (currentUser) {
    el.innerHTML = `
      <div style="display:flex;align-items:center;gap:8px">
        <div class="user-pill" onclick="openDbPanel()">
          <div class="user-avatar">${currentUser.nickname.charAt(0).toUpperCase()}</div>
          <span class="user-name">${currentUser.nickname}</span>
          <span style="font-size:10px;color:#aaa;margin-left:2px">🗄️</span>
        </div>
        <button class="logout-btn" onclick="doLogout()">로그아웃</button>
      </div>`;
  } else {
    el.innerHTML = `
      <button class="login-btn outline" onclick="openAuth('login')">로그인</button>
      <button class="login-btn" onclick="openAuth('signup')">회원가입</button>`;
  }
}

function openAuth(tab) { document.getElementById('authOverlay').classList.add('open'); switchTab(tab); clearAuthError(); }
function closeAuth() { document.getElementById('authOverlay').classList.remove('open'); }
function closeAuthIfOverlay(e) { if (e.target === document.getElementById('authOverlay')) closeAuth(); }
function clearAuthError() { const el = document.getElementById('authError'); el.classList.remove('show'); el.textContent = ''; }
function showAuthError(msg) { const el = document.getElementById('authError'); el.textContent = msg; el.classList.add('show'); }

function switchTab(tab) {
  document.getElementById('tabLogin').classList.toggle('active', tab === 'login');
  document.getElementById('tabSignup').classList.toggle('active', tab === 'signup');
  document.getElementById('formLogin').style.display = tab === 'login' ? 'block' : 'none';
  document.getElementById('formSignup').style.display = tab === 'signup' ? 'block' : 'none';
  document.getElementById('authSub').textContent = tab === 'login' ? '반갑습니다! 로그인해 주세요' : '지금 가입하고 레시피를 저장해 보세요';
  clearAuthError();
}

async function doSignup() {
  clearAuthError();
  const nick = document.getElementById('signupNick').value.trim();
  const email = document.getElementById('signupEmail').value.trim();
  const pw = document.getElementById('signupPw').value;
  const pw2 = document.getElementById('signupPw2').value;
  if (!nick) { showAuthError('닉네임을 입력해주세요.'); return; }
  if (!email.includes('@')) { showAuthError('올바른 이메일을 입력해주세요.'); return; }
  if (pw.length < 6) { showAuthError('비밀번호는 6자 이상이어야 합니다.'); return; }
  if (pw !== pw2) { showAuthError('비밀번호가 일치하지 않습니다.'); return; }
  await openDB();
  const existing = await dbGetByEmail(email);
  if (existing) { showAuthError('이미 사용 중인 이메일입니다.'); return; }
  const user = { email, nickname: nick, password: simpleHash(pw), createdAt: new Date().toISOString() };
  const id = await dbAdd('users', user);
  user.id = id;
  closeAuth();
  setUser({ id, email, nickname: nick });
  showToast('🎉 ' + nick + '님, 환영합니다!', 'success');
}

async function doLogin() {
  clearAuthError();
  const email = document.getElementById('loginEmail').value.trim();
  const pw = document.getElementById('loginPw').value;
  if (!email || !pw) { showAuthError('이메일과 비밀번호를 모두 입력해주세요.'); return; }
  await openDB();
  const user = await dbGetByEmail(email);
  if (!user) { showAuthError('존재하지 않는 계정입니다.'); return; }
  if (user.password !== simpleHash(pw)) { showAuthError('비밀번호가 틀렸습니다.'); return; }
  closeAuth();
  setUser({ id: user.id, email: user.email, nickname: user.nickname });
  showToast('👋 ' + user.nickname + '님, 다시 돌아오셨군요!', 'success');
}

function doLogout() {
  if (!confirm('로그아웃 하시겠습니까?')) return;
  currentUser = null;
  sessionStorage.removeItem('naengUser');
  ingredients = [];
  expiryItems = [];
  renderHeaderAuth();
  renderChips();
  showToast('로그아웃 되었습니다.', 'warn');
}

async function loadUserData() {
  if (!currentUser) { ingredients = []; expiryItems = []; renderChips(); return; }
  await openDB();
  const rows = await dbGetByIndex('ingredients', 'userId', currentUser.id);
  ingredients = rows.map(r => r.name);
  const exp = await dbGetByIndex('expiry_items', 'userId', currentUser.id);
  expiryItems = exp;
  renderChips();
  if (document.getElementById('page-calendar').classList.contains('active')) renderCalendar();
}

let activeDbTab = 'users';
function openDbPanel() { document.getElementById('dbOverlay').classList.add('open'); switchDbTab('users'); }
function closeDbPanel() { document.getElementById('dbOverlay').classList.remove('open'); }
function closeDbIfOverlay(e) { if (e.target === document.getElementById('dbOverlay')) closeDbPanel(); }

async function switchDbTab(tab) {
  activeDbTab = tab;
  document.querySelectorAll('.db-tab').forEach(t => t.classList.toggle('active', t.textContent.includes(tab.split('_')[0])));
  await openDB();
  const body = document.getElementById('dbBody');
  if (tab === 'sql') {
    body.innerHTML = `
      <div style="color:#7dd3fc;font-size:13px;font-weight:700;margin-bottom:12px">SQL-Like 쿼리 (IndexedDB 시뮬레이션)</div>
      <div class="sql-box">
        <textarea id="sqlInput" placeholder="예시: SELECT * FROM users\nSELECT * FROM ingredients WHERE userId = 1\nDELETE FROM expiry_items WHERE id = 3\nCOUNT users"></textarea>
      </div>
      <button class="sql-run" onclick="runSqlQuery()">▶ 실행</button>
      <div class="sql-result" id="sqlResult"></div>`;
    return;
  }
  if (tab === 'users') {
    const users = await dbGetAll('users');
    const stats = `<div class="db-stat-row">
      <div class="db-stat"><div class="db-stat-n">${users.length}</div><div class="db-stat-l">총 가입자 수</div></div>
      <div class="db-stat"><div class="db-stat-n">${currentUser ? 1 : 0}</div><div class="db-stat-l">현재 세션</div></div>
    </div>`;
    if (!users.length) { body.innerHTML = stats + '<div class="db-empty">등록된 사용자가 없습니다.</div>'; return; }
    body.innerHTML = stats + `<table class="db-table">
      <thead><tr><th>ID</th><th>닉네임</th><th>이메일</th><th>가입일</th><th>비밀번호(해시)</th></tr></thead>
      <tbody>${users.map(u => `<tr>
        <td><span class="db-badge gray">${u.id}</span></td>
        <td><b style="color:#7dd3fc">${u.nickname}</b></td>
        <td>${u.email}</td>
        <td style="color:#888">${u.createdAt ? u.createdAt.slice(0,10) : '-'}</td>
        <td style="font-size:11px;font-family:monospace;color:#666">${u.password}</td>
      </tr>`).join('')}</tbody>
    </table>`;
  }
  if (tab === 'ingredients') {
    const rows = await dbGetAll('ingredients');
    const stats = `<div class="db-stat-row">
      <div class="db-stat"><div class="db-stat-n">${rows.length}</div><div class="db-stat-l">총 재료 레코드</div></div>
      <div class="db-stat"><div class="db-stat-n">${[...new Set(rows.map(r => r.userId))].length}</div><div class="db-stat-l">등록 사용자 수</div></div>
    </div>`;
    if (!rows.length) { body.innerHTML = stats + '<div class="db-empty">등록된 재료가 없습니다.</div>'; return; }
    body.innerHTML = stats + `<table class="db-table">
      <thead><tr><th>ID</th><th>UserID</th><th>재료명</th><th>등록일</th></tr></thead>
      <tbody>${rows.map(r => `<tr>
        <td><span class="db-badge gray">${r.id}</span></td>
        <td><span class="db-badge green">${r.userId}</span></td>
        <td>${r.name}</td>
        <td style="color:#888">${r.createdAt ? r.createdAt.slice(0,10) : '-'}</td>
      </tr>`).join('')}</tbody>
    </table>`;
  }
  if (tab === 'expiry') {
    const rows = await dbGetAll('expiry_items');
    const today = new Date(); today.setHours(0,0,0,0);
    const expired = rows.filter(r => new Date(r.expiry) < today).length;
    const stats = `<div class="db-stat-row">
      <div class="db-stat"><div class="db-stat-n">${rows.length}</div><div class="db-stat-l">총 유통기한 항목</div></div>
      <div class="db-stat"><div class="db-stat-n" style="color:#ef4444">${expired}</div><div class="db-stat-l">만료된 항목</div></div>
    </div>`;
    if (!rows.length) { body.innerHTML = stats + '<div class="db-empty">등록된 유통기한 항목이 없습니다.</div>'; return; }
    body.innerHTML = stats + `<table class="db-table">
      <thead><tr><th>ID</th><th>UserID</th><th>재료명</th><th>유통기한</th><th>자동등록</th><th>등록일</th></tr></thead>
      <tbody>${rows.map(r => `<tr>
        <td><span class="db-badge gray">${r.id}</span></td>
        <td><span class="db-badge green">${r.userId}</span></td>
        <td>${r.name}</td>
        <td style="color:${new Date(r.expiry) < today ? '#ef4444' : '#6ee7b7'}">${r.expiry}</td>
        <td>${r.autoAdded ? '<span class="db-badge green">자동</span>' : '-'}</td>
        <td style="color:#888">${r.createdAt ? r.createdAt.slice(0,10) : '-'}</td>
      </tr>`).join('')}</tbody>
    </table>`;
  }
}

async function runSqlQuery() {
  const raw = document.getElementById('sqlInput').value.trim().toUpperCase();
  const result = document.getElementById('sqlResult');
  await openDB();
  try {
    if (raw.startsWith('SELECT * FROM')) {
      const parts = raw.split(' ');
      const table = parts[3].toLowerCase();
      const storeMap = { users:'users', ingredients:'ingredients', expiry_items:'expiry_items' };
      const st = storeMap[table];
      if (!st) { result.textContent = '❌ 알 수 없는 테이블: ' + parts[3]; return; }
      let rows = await dbGetAll(st);
      const whereIdx = raw.indexOf('WHERE USERID =');
      if (whereIdx > -1) {
        const uid = parseInt(raw.split('WHERE USERID =')[1].trim());
        rows = rows.filter(r => r.userId === uid);
      }
      result.textContent = JSON.stringify(rows, null, 2);
    } else if (raw.startsWith('COUNT')) {
      const table = raw.split(' ')[1].toLowerCase();
      const storeMap = { users:'users', ingredients:'ingredients', expiry_items:'expiry_items' };
      const st = storeMap[table];
      if (!st) { result.textContent = '❌ 알 수 없는 테이블.'; return; }
      const rows = await dbGetAll(st);
      result.textContent = `COUNT(*) → ${rows.length}개`;
    } else if (raw.startsWith('DELETE FROM')) {
      const parts = raw.split(' ');
      const table = parts[2].toLowerCase();
      const storeMap = { users:'users', ingredients:'ingredients', expiry_items:'expiry_items' };
      const st = storeMap[table];
      if (!st) { result.textContent = '❌ 알 수 없는 테이블.'; return; }
      const whereMatch = raw.match(/WHERE ID = (\d+)/);
      if (!whereMatch) { result.textContent = '❌ WHERE id = N 절이 필요합니다.'; return; }
      await dbDelete(st, parseInt(whereMatch[1]));
      result.textContent = `✅ ${table}에서 id=${whereMatch[1]} 삭제 완료`;
      await loadUserData();
    } else {
      result.textContent = '지원 문법:\nSELECT * FROM users\nSELECT * FROM ingredients WHERE userId = 1\nSELECT * FROM expiry_items\nCOUNT users\nDELETE FROM expiry_items WHERE id = 3';
    }
  } catch (e) {
    result.textContent = '❌ 오류: ' + e.message;
  }
}

let ingredients = [];

async function addIngredient() {
  const input = document.getElementById('ingredientInput');
  const val = input.value.trim();
  if (!val) return;
  if (currentUser) {
    await openDB();
    const existing = await dbGetByIndex('ingredients', 'userId', currentUser.id);
    if (existing.find(r => r.name === val)) { input.value = ''; return; }
    await dbAdd('ingredients', { userId: currentUser.id, name: val, createdAt: new Date().toISOString() });
    ingredients.push(val);
  } else {
    if (!ingredients.includes(val)) ingredients.push(val);
  }
  input.value = '';
  renderChips();
  renderIngredients();
}

async function removeIngredient(i) {
  const name = ingredients[i];
  if (currentUser) {
    await openDB();
    const rows = await dbGetByIndex('ingredients', 'userId', currentUser.id);
    const found = rows.find(r => r.name === name);
    if (found) await dbDelete('ingredients', found.id);
  }
  ingredients.splice(i, 1);
  renderChips();
}

function renderChips() {
  const c = document.getElementById('chipContainer');
  if (!ingredients.length) { c.innerHTML = '<span style="color:#aaa;font-size:13px;">등록된 재료가 없습니다.</span>'; return; }
  c.innerHTML = ingredients.map((ing, i) =>
    `<span class="chip"># ${ing} <span class="chip-remove" onclick="removeIngredient(${i})">×</span></span>`
  ).join('');
}

const PRESET_RECIPES = [
  { RCP_NM:'계란볶음밥', INFO_ENG:'450', RCP_PARTS_DTLS:'계란, 밥, 대파, 간장, 식용유', ATT_FILE_NO_MAIN:'https://images.unsplash.com/photo-1512058564366-18510be2db19?q=80&w=800', MANUAL01:'대파를 쫑쫑 썰어 파기름을 냅니다.', MANUAL02:'밥과 계란을 넣고 고슬고슬하게 볶습니다.', MANUAL03:'간장 한 스푼으로 불맛을 내어 완성합니다.' },
  { RCP_NM:'참치김치찌개', INFO_ENG:'320', RCP_PARTS_DTLS:'김치, 참치캔, 두부, 대파, 고춧가루', ATT_FILE_NO_MAIN:'https://images.unsplash.com/photo-1585032226651-759b368d7246?q=80&w=800', MANUAL01:'냄비에 김치와 참치기름을 넣고 달달 볶습니다.', MANUAL02:'물을 붓고 끓으면 참치와 두부를 넣습니다.', MANUAL03:'한소끔 끓여 대파를 올려 마무리합니다.' },
  { RCP_NM:'크림파스타', INFO_ENG:'620', RCP_PARTS_DTLS:'파스타면, 생크림, 베이컨, 양파, 마늘', ATT_FILE_NO_MAIN:'https://images.unsplash.com/photo-1555949258-eb67b1ef0ceb?q=80&w=800', MANUAL01:'면을 소금물에 7분간 삶아냅니다.', MANUAL02:'팬에 마늘, 양파, 베이컨을 볶다가 생크림을 붓습니다.', MANUAL03:'소스가 끓으면 면을 넣고 졸여 완성합니다.' },
  { RCP_NM:'두부김치치즈구이', INFO_ENG:'280', RCP_PARTS_DTLS:'두부, 신김치, 모짜렐라치즈, 참기름', ATT_FILE_NO_MAIN:'https://images.unsplash.com/photo-1546069901-ba9599a7e63c?q=80&w=400', MANUAL01:'두부를 먹기 좋은 크기로 썰어 수분을 뺍니다.', MANUAL02:'볶은 김치를 두부 위에 올리고 치즈를 뿌립니다.', MANUAL03:'에어프라이어에서 치즈가 녹을 때까지 돌립니다.' },
  { RCP_NM:'소고기무국', INFO_ENG:'190', RCP_PARTS_DTLS:'소고기, 무, 대파, 국간장, 다진마늘', ATT_FILE_NO_MAIN:'https://images.unsplash.com/photo-1547592180-85f173990554?q=80&w=1200', MANUAL01:'참기름을 두른 냄비에 소고기와 무를 볶습니다.', MANUAL02:'물을 붓고 거품을 걷어내며 푹 끓입니다.', MANUAL03:'다진마늘과 국간장으로 간을 하고 파를 넣습니다.' }
];

const RECIPE_ICONS = ['🍳','🥘','🍜','🥗','🍲','🥩','🍱','🍛','🥚','🍝'];

function renderMainAndAllRecipes() {
  const pc = document.getElementById('popularCardsContainer');
  pc.innerHTML = PRESET_RECIPES.slice(0,3).map((r, i) => `
    <div class="recipe-card-noimg" onclick="openModalFromPreset('${r.RCP_NM}')">
      <div class="recipe-card-noimg-body">
        <div class="recipe-card-noimg-icon">${RECIPE_ICONS[i % RECIPE_ICONS.length]}</div>
        <h3>${r.RCP_NM}</h3>
        <p>${r.RCP_PARTS_DTLS.substring(0,30)}...</p>
        <div class="calorie-info" style="margin-top:8px;">🔥 ${r.INFO_ENG} kcal</div>
      </div>
    </div>`).join('');
  renderAllRecipesTable([...PRESET_RECIPES].sort((a,b) => Number(a.INFO_ENG)-Number(b.INFO_ENG)));
  updateRecipeCount(PRESET_RECIPES.length, PRESET_RECIPES.length);
}

function renderAllRecipesTable(list) {
  const ac = document.getElementById('allRecipesCardsContainer');
  if (!list.length) { ac.innerHTML = '<div class="no-recipe-result">😢 검색 결과가 없습니다.</div>'; return; }
  ac.innerHTML = list.map((r, i) => `
    <div class="recipe-row" onclick="openModalFromPreset('${r.RCP_NM}')">
      <div class="recipe-row-num">${RECIPE_ICONS[i % RECIPE_ICONS.length]}</div>
      <div class="recipe-row-info">
        <div class="recipe-row-name">${r.RCP_NM}</div>
        <div class="recipe-row-parts">${r.RCP_PARTS_DTLS}</div>
      </div>
      <div class="recipe-row-kcal">🔥 ${r.INFO_ENG} kcal</div>
      <div class="recipe-row-arrow">›</div>
    </div>`).join('');
}

function updateRecipeCount(shown, total) {
  const el = document.getElementById('recipeSearchCount');
  if (!el) return;
  el.textContent = shown === total ? `총 ${total}개 레시피` : `${shown}개 검색됨 (전체 ${total}개)`;
}

function filterRecipes() {
  const q = document.getElementById('recipeSearchInput').value.trim().toLowerCase();
  const sorted = [...PRESET_RECIPES].sort((a,b) => Number(a.INFO_ENG)-Number(b.INFO_ENG));
  if (!q) { renderAllRecipesTable(sorted); updateRecipeCount(sorted.length, PRESET_RECIPES.length); return; }
  const filtered = sorted.filter(r => r.RCP_NM.toLowerCase().includes(q) || r.RCP_PARTS_DTLS.toLowerCase().includes(q));
  renderAllRecipesTable(filtered);
  updateRecipeCount(filtered.length, PRESET_RECIPES.length);
}

function openModalFromPreset(name) {
  const f = PRESET_RECIPES.find(r => r.RCP_NM === name);
  if (f) openModal(f);
}

const EXPIRY_DB = { '소고기':3,'돼지고기':3,'닭고기':2,'닭':2,'삼겹살':3,'다짐육':2,'베이컨':7,'햄':7,'소시지':7,'스팸':7,'생선':2,'고등어':2,'삼치':2,'연어':2,'참치':3,'새우':2,'오징어':2,'꽃게':2,'조개':2,'굴':3,'계란':21,'달걀':21,'우유':7,'치즈':14,'버터':30,'요거트':14,'생크림':5,'두부':4,'김치':30,'시금치':4,'상추':4,'양상추':5,'깻잎':4,'부추':4,'대파':7,'파':7,'양파':30,'마늘':30,'생강':14,'고추':7,'당근':14,'감자':30,'고구마':30,'무':14,'배추':14,'브로콜리':5,'애호박':5,'가지':5,'오이':5,'토마토':7,'파프리카':7,'버섯':5,'밥':1,'된장':180,'고추장':180,'간장':365 };

function getExpiryDays(name) {
  if (EXPIRY_DB[name] !== undefined) return EXPIRY_DB[name];
  for (const key of Object.keys(EXPIRY_DB)) {
    if (name.includes(key) || key.includes(name)) return EXPIRY_DB[key];
  }
  return null;
}

function calcExpiryDate(days) {
  const d = new Date();
  d.setDate(d.getDate() + days);
  return d.getFullYear() + '-' + String(d.getMonth()+1).padStart(2,'0') + '-' + String(d.getDate()).padStart(2,'0');
}

async function autoRegisterIngredients(list) {
  if (!currentUser) return;
  const added = [], skipped = [];
  for (const ing of list) {
    const days = getExpiryDays(ing);
    if (days === null) { skipped.push(ing); continue; }
    if (expiryItems.some(it => it.name === ing)) continue;
    const item = { userId: currentUser.id, name: ing, expiry: calcExpiryDate(days), autoAdded: true, createdAt: new Date().toISOString() };
    const id = await dbAdd('expiry_items', item);
    item.id = id;
    expiryItems.push(item);
    added.push(ing + '(' + days + '일)');
  }
  if (added.length > 0) showToast('📅 캘린더 자동 등록: ' + added.join(', '), 'success');
  if (skipped.length > 0) setTimeout(() => showToast('⚠️ 유통기한 미확인: ' + skipped.join(', '), 'warn'), 1500);
}

let recipeList = [], recipeNames = [], currentAngle = 0, isSpinning = false;
const COLORS = ['#34d399','#6ee7b7','#a7f3d0','#10b981','#059669','#d1fae5'];
const START_OFFSET = -Math.PI / 2;

async function fetchRecipesAndGo() {
  if (ingredients.length === 0) { alert('냉장고에 재료를 최소 한 개 이상 등록해주세요!'); return; }
  document.getElementById('loadingOverlay').classList.add('show');
  recipeList = [];
  try {
    const res = await fetch('http://localhost:8080/api/mix-recipe', {
      method: 'POST', headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ ingredients })
    });
    if (!res.ok) throw new Error();
    const data = await res.json();
    if (data.COOKRCP01?.row) {
      recipeList = data.COOKRCP01.row.filter(r => ingredients.some(ing => r.RCP_NM.includes(ing)));
    }
  } catch (e) { recipeList = []; }
  document.getElementById('loadingOverlay').classList.remove('show');
  document.getElementById('loadingOverlay').style.display = 'none';
  if (recipeList.length < 2) {
    recipeList = [];
    ingredients.forEach(ing => {
      recipeList.push(
        { RCP_NM: ing + ' 볶음밥', INFO_ENG:'430', RCP_PARTS_DTLS: ing + ', 밥, 계란, 소금', ATT_FILE_NO_MAIN:'https://images.unsplash.com/photo-1512058564366-18510be2db19?q=80&w=800', MANUAL01:'재료를 잘게 준비합니다.', MANUAL02:'팬에 기름을 두르고 달굽니다.', MANUAL03: ing + '와 밥을 넣고 볶습니다.' },
        { RCP_NM: ing + ' 전골', INFO_ENG:'310', RCP_PARTS_DTLS: ing + ', 대파, 두부, 양념장', ATT_FILE_NO_MAIN:'https://images.unsplash.com/photo-1546069901-ba9599a7e63c?q=80&w=400', MANUAL01:'재료를 썰어 전골 냄비에 담습니다.', MANUAL02:'육수를 붓고 끓입니다.', MANUAL03:'간을 맞춰 완성합니다.' }
      );
    });
  }
  if (recipeList.length > 6) recipeList = recipeList.slice(0, 6);
  recipeNames = recipeList.map(r => r.RCP_NM);
  await autoRegisterIngredients(ingredients);
  showPage('roulette', null);
  currentAngle = 0;
  renderSideList();
  drawRoulette();
}

function drawRoulette() {
  const canvas = document.getElementById('rouletteCanvas');
  if (!canvas) return;
  const ctx = canvas.getContext('2d');
  const n = recipeNames.length;
  if (!n) return;
  const seg = (2*Math.PI)/n;
  const cx = 170, cy = 170, r = 165;
  ctx.clearRect(0,0,340,340);
  recipeNames.forEach((name, i) => {
    const s = START_OFFSET + currentAngle + i*seg, en = s + seg;
    ctx.beginPath(); ctx.moveTo(cx,cy); ctx.arc(cx,cy,r,s,en);
    ctx.fillStyle = COLORS[i % COLORS.length]; ctx.fill();
    ctx.strokeStyle = '#fff'; ctx.lineWidth = 2; ctx.stroke();
    ctx.save(); ctx.translate(cx,cy); ctx.rotate(s + seg/2);
    ctx.textAlign = 'right'; ctx.fillStyle = '#064e3b'; ctx.font = 'bold 13px Noto Sans KR';
    ctx.fillText(name.length > 9 ? name.substring(0,9)+'…' : name, r-12, 5);
    ctx.restore();
  });
  ctx.beginPath(); ctx.arc(cx,cy,22,0,2*Math.PI); ctx.fillStyle='#fff'; ctx.fill();
  ctx.strokeStyle='#e5e7eb'; ctx.lineWidth=2; ctx.stroke();
}

function spinRoulette() {
  if (isSpinning) return;
  isSpinning = true;
  document.getElementById('spinBtn').disabled = true;
  const totalRot = (5 + Math.floor(Math.random()*5))*2*Math.PI + Math.random()*2*Math.PI;
  const duration = 3500;
  let startTime = null, startAngle = currentAngle;
  function animate(ts) {
    if (!startTime) startTime = ts;
    const progress = Math.min((ts-startTime)/duration, 1);
    const ease = 1 - Math.pow(1-progress, 4);
    currentAngle = startAngle + ease * totalRot;
    drawRoulette();
    if (progress < 1) { requestAnimationFrame(animate); return; }
    isSpinning = false;
    document.getElementById('spinBtn').disabled = false;
    const n = recipeNames.length;
    const raw = ((-currentAngle) % (2*Math.PI) + 2*Math.PI) % (2*Math.PI);
    const winIdx = Math.floor(raw / ((2*Math.PI)/n)) % n;
    document.querySelectorAll('.side-card').forEach((el,i) => el.classList.toggle('active', i===winIdx));
    openModal(recipeList[winIdx]);
  }
  requestAnimationFrame(animate);
}

function renderSideList() {
  const list = document.getElementById('sideList');
  document.getElementById('recipeCount').textContent = recipeList.length;
  list.innerHTML = recipeList.map((r, i) =>
    `<div class="side-card" onclick="openModal(recipeList[${i}]);highlightSide(${i})">
      <div class="side-name">${r.RCP_NM}</div>
      <div class="side-parts">${r.RCP_PARTS_DTLS || '재료 정보 없음'}</div>
    </div>`
  ).join('');
}

function highlightSide(i) { document.querySelectorAll('.side-card').forEach((el,idx) => el.classList.toggle('active', idx===i)); }

function resetAll() {
  ingredients = []; recipeList = []; recipeNames = []; currentAngle = 0;
  renderChips();
  showPage('home', document.getElementById('menu-home'));
}

function openModal(recipe) {
  document.getElementById('modalTitle').textContent = recipe.RCP_NM;
  document.getElementById('modalCalorie').textContent = `🔥 ${recipe.INFO_ENG || '0'} kcal`;
  document.getElementById('modalIngredients').textContent = `재료: ${recipe.RCP_PARTS_DTLS || '정보 없음'}`;
  const sc = document.getElementById('modalSteps'); sc.innerHTML = '';
  for (let i = 1; i <= 20; i++) {
    const key = i < 10 ? 'MANUAL0'+i : 'MANUAL'+i;
    if (recipe[key]?.trim()) {
      const div = document.createElement('div'); div.className = 'step';
      div.innerHTML = `<span class="step-num">${i}.</span><span>${recipe[key].replace(/\\n/g,' ')}</span>`;
      sc.appendChild(div);
    }
  }
  document.getElementById('recipeModal').classList.add('open');
}
function closeModal() { document.getElementById('recipeModal').classList.remove('open'); }
function closeModalIfOverlay(e) { if (e.target === document.getElementById('recipeModal')) closeModal(); }

let expiryItems = [], calYear, calMonth;

function initCalendar() { const n=new Date(); calYear=n.getFullYear(); calMonth=n.getMonth(); }
function changeMonth(d) {
  calMonth += d;
  if (calMonth>11){calMonth=0;calYear++;} if (calMonth<0){calMonth=11;calYear--;}
  renderCalendar();
}
function renderCalendar() {
  checkExpiryAlerts();
  const label = document.getElementById('calMonthLabel'); if(!label) return;
  label.textContent = `${calYear}년 ${calMonth+1}월`;
  const grid = document.getElementById('calGrid'); grid.innerHTML = '';
  ['일','월','화','수','목','금','토'].forEach(d => { const el=document.createElement('div'); el.className='cal-day-label'; el.textContent=d; grid.appendChild(el); });
  const firstDay = new Date(calYear,calMonth,1).getDay();
  const dim = new Date(calYear,calMonth+1,0).getDate();
  const today = new Date();
  for (let i=0;i<firstDay;i++) { const el=document.createElement('div'); el.className='cal-cell other-month'; grid.appendChild(el); }
  for (let d=1;d<=dim;d++) {
    const dateStr = `${calYear}-${String(calMonth+1).padStart(2,'0')}-${String(d).padStart(2,'0')}`;
    const items = expiryItems.filter(it=>it.expiry===dateStr);
    const isToday = today.getFullYear()===calYear && today.getMonth()===calMonth && today.getDate()===d;
    const el = document.createElement('div');
    el.className = 'cal-cell'+(isToday?' today':'')+(items.length?' has-items':'');
    el.innerHTML = `<span>${d}</span>`;
    if (items.length) {
      const dr=document.createElement('div'); dr.className='dot-row';
      items.forEach(it => { const dot=document.createElement('div'); dot.className='expiry-dot'; dot.style.background=getStatusColor(it.expiry); dr.appendChild(dot); });
      el.appendChild(dr);
    }
    grid.appendChild(el);
  }
  renderItemsList();
}

function getStatusColor(e) { const d=new Date(); d.setHours(0,0,0,0); const diff=Math.ceil((new Date(e)-d)/86400000); if(diff<0)return'#9ca3af'; if(diff<=2)return'#ef4444'; if(diff<=5)return'#f59e0b'; return'#2bb673'; }
function getStatusClass(e) { const d=new Date(); d.setHours(0,0,0,0); const diff=Math.ceil((new Date(e)-d)/86400000); if(diff<0)return'expired'; if(diff<=2)return'danger'; if(diff<=5)return'warn'; return'safe'; }
function getStatusLabel(e) { const d=new Date(); d.setHours(0,0,0,0); const diff=Math.ceil((new Date(e)-d)/86400000); if(diff<0)return'만료됨'; if(diff===0)return'D-Day'; return`D-${diff}`; }

async function addExpiryItem() {
  const name = document.getElementById('itemName').value.trim();
  const expiry = document.getElementById('itemExpiry').value;
  if (!name || !expiry) { alert('재료 이름과 유통기한을 모두 입력해주세요.'); return; }
  const item = { name, expiry, autoAdded: false, createdAt: new Date().toISOString() };
  if (currentUser) {
    await openDB();
    item.userId = currentUser.id;
    const id = await dbAdd('expiry_items', item);
    item.id = id;
  } else {
    item.id = Date.now() + Math.random();
  }
  expiryItems.push(item);
  document.getElementById('itemName').value = '';
  document.getElementById('itemExpiry').value = '';
  renderCalendar();
}

async function removeExpiryItem(id) {
  expiryItems = expiryItems.filter(it => it.id !== id);
  if (currentUser) { await openDB(); await dbDelete('expiry_items', id); }
  renderCalendar();
}

function renderItemsList() {
  const c = document.getElementById('itemsList'); if (!c) return;
  if (!expiryItems.length) { c.innerHTML = '<div style="color:#aaa;font-size:13px;text-align:center;padding:16px 0;">등록된 재료가 없습니다.</div>'; return; }
  c.innerHTML = [...expiryItems].sort((a,b)=>a.expiry.localeCompare(b.expiry)).map(it =>
    `<div class="expiry-item ${getStatusClass(it.expiry)}">
      <div style="flex:1">
        <div class="item-name">${it.name}${it.autoAdded?'<span style="font-size:10px;background:#e8fff3;color:#1a9057;padding:2px 7px;border-radius:999px;font-weight:700;margin-left:6px;">자동</span>':''}</div>
        <div class="item-date">유통기한: ${it.expiry}</div>
      </div>
      <span class="item-badge">${getStatusLabel(it.expiry)}</span>
      <button class="item-del" onclick="removeExpiryItem(${it.id})">🗑</button>
    </div>`
  ).join('');
}

function checkExpiryAlerts() {
  const today = new Date(); today.setHours(0,0,0,0);
  const urgent = expiryItems.filter(it => { const diff=Math.ceil((new Date(it.expiry)-today)/86400000); return diff>=0&&diff<=3; });
  const banner = document.getElementById('alertBanner'); if (!banner) return;
  if (!urgent.length) { banner.classList.add('hidden'); return; }
  banner.classList.remove('hidden');
  document.getElementById('alertTitle').textContent = `⚠️ 유통기한 임박 재료 ${urgent.length}개!`;
  document.getElementById('alertBody').textContent = urgent.map(it => {
    const diff = Math.ceil((new Date(it.expiry)-today)/86400000);
    return `${it.name}(${diff===0?'오늘 마감':diff+'일 남음'})`;
  }).join(' · ');
}

function showPage(name, el) {
  document.querySelectorAll('.page').forEach(p => p.classList.remove('active'));
  document.getElementById('page-' + name).classList.add('active');
  document.querySelectorAll('.nav-menu').forEach(a => a.classList.remove('active'));
  if (el) el.classList.add('active');
  else {
    if (name==='home') document.getElementById('menu-home').classList.add('active');
    if (name==='calendar') document.getElementById('menu-calendar').classList.add('active');
  }
  if (name === 'calendar') renderCalendar();
  if (name === 'recipe-list-page') {
    const si = document.getElementById('recipeSearchInput');
    if (si) { si.value = ''; filterRecipes(); }
  }
}

function showToast(msg, type) {
  const old = document.getElementById('toastMsg'); if (old) old.remove();
  const t = document.createElement('div'); t.id='toastMsg';
  t.style.cssText = 'position:fixed;bottom:28px;left:50%;transform:translateX(-50%);background:'+(type==='success'?'#1a9057':'#d97706')+';color:#fff;padding:14px 26px;border-radius:14px;font-size:14px;font-weight:600;font-family:Noto Sans KR,sans-serif;box-shadow:0 8px 24px rgba(0,0,0,.2);z-index:9999;max-width:500px;text-align:center;line-height:1.5;';
  t.textContent = msg; document.body.appendChild(t);
  setTimeout(()=>{ t.style.transition='opacity .4s'; t.style.opacity='0'; setTimeout(()=>t.remove(),400); }, 3800);
}

document.addEventListener('DOMContentLoaded', function() {
  const si = document.getElementById('recipeSearchInput');
  if (si) {
    si.addEventListener('focus', () => si.style.borderColor = 'var(--green)');
    si.addEventListener('blur', () => si.style.borderColor = '#e5e7eb');
  }
});

window.onload = async function() {
  initCalendar();
  renderMainAndAllRecipes();
  renderHeaderAuth();
  await openDB();
  if (currentUser) await loadUserData();
  else renderChips();
};
</script>
</body>
</html>
