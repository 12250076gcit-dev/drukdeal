// Load admin dashboard data
async function loadAdminDashboard() {
  try {
    // Fetch metrics from backend
    const metricsRes = await fetch('/api/admin/metrics', { cache: 'no-store' });
    const metricsData = await metricsRes.json();

    // Update stat cards with metrics
    const totalUsersEl = document.getElementById('totalUsersValue');
    const totalProductsEl = document.getElementById('totalProductsValue');
    
    if (totalUsersEl) {
      totalUsersEl.textContent = formatNumber(metricsData.total_users || 0);
    }
    if (totalProductsEl) {
      totalProductsEl.textContent = formatNumber(metricsData.total_products || 0);
    }

    // Fetch products from backend
    const productsRes = await fetch('/products', { cache: 'no-store' });
    const products = await productsRes.json();

    // Update pending count
    const pendingBadgeEl = document.getElementById('pendingBadge');
    const statuses = JSON.parse(localStorage.getItem('admin_product_statuses') || '{}');
    const pendingCount = Object.values(statuses).filter(s => s === 'pending').length || products.length;
    if (pendingBadgeEl) {
      pendingBadgeEl.textContent = `Pending ${pendingCount}`;
    }

    // Render products list
    renderProductList(products);
  } catch (err) {
    console.error('Failed to load admin dashboard:', err);
    const container = document.getElementById('listingListContainer');
    if (container) {
      container.innerHTML = '<div style="padding: 20px; text-align: center; color: #888;">Failed to load data</div>';
    }
  }
}

function renderProductList(products) {
  const container = document.getElementById('listingListContainer');
  if (!container) return;

  const statuses = JSON.parse(localStorage.getItem('admin_product_statuses') || '{}');
  const visibleProducts = products.filter(product => statuses[product.prod_id] !== 'rejected');
  const itemsToShow = visibleProducts.slice(0, 8);

  if (itemsToShow.length === 0) {
    container.innerHTML = '<div style="padding: 20px; text-align: center; color: #888;">No products available</div>';
    return;
  }

  container.innerHTML = itemsToShow.map(product => {
    const productStatus = statuses[product.prod_id] || 'pending';
    const displayImage = product.image_path
      ? `<img src="${product.image_path}" alt="${product.title}" class="listing-thumb" style="object-fit: cover;" onerror="this.style.display='none'" />`
      : '<div class="thumb-placeholder"></div>';

    return `
      <div class="listing-item">
        ${displayImage}
        <div style="flex: 1;">
          <div class="listing-name">${product.title}</div>
          <div style="font-size: 0.8rem; color: #888;">Nu. ${formatPrice(product.price)} • ${product.seller_name}</div>
          <div style="font-size: 0.75rem; color: #aaa; margin-top: 4px;">${product.created_at}</div>
        </div>
        <div class="listing-actions">
          <button class="btn btn-accept" onclick="handleApprove(${product.prod_id})">
            ${productStatus === 'approved' ? '✓ Approved' : 'Accept'}
          </button>
          <button class="btn btn-reject" onclick="handleReject(${product.prod_id})">
            ${productStatus === 'rejected' ? '✗ Rejected' : 'Reject'}
          </button>
        </div>
      </div>
    `;
  }).join('');
}

function handleApprove(prodId) {
  const statuses = JSON.parse(localStorage.getItem('admin_product_statuses') || '{}');
  statuses[prodId] = 'approved';
  localStorage.setItem('admin_product_statuses', JSON.stringify(statuses));
  loadAdminDashboard();
}

function handleReject(prodId) {
  const statuses = JSON.parse(localStorage.getItem('admin_product_statuses') || '{}');
  statuses[prodId] = 'rejected';
  localStorage.setItem('admin_product_statuses', JSON.stringify(statuses));
  loadAdminDashboard();
}

function formatNumber(num) {
  return Math.abs(num) > 999
    ? Math.sign(num) * ((Math.abs(num) / 1000).toFixed(1)) + 'k'
    : num.toLocaleString();
}

function formatPrice(price) {
  return parseFloat(price || 0).toLocaleString('en-IN', {
    minimumFractionDigits: 0,
    maximumFractionDigits: 2
  });
}

// Load dashboard on page load
document.addEventListener('DOMContentLoaded', loadAdminDashboard);

// Reload every 30 seconds
setInterval(loadAdminDashboard, 30000);
