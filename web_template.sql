-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: mariadb
-- Generation Time: Aug 29, 2023 at 07:58 AM
-- Server version: 10.11.2-MariaDB-1:10.11.2+maria~ubu2204
-- PHP Version: 8.1.17

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `web_0`
--

-- --------------------------------------------------------

--
-- Table structure for table `addresses`
--

CREATE TABLE `addresses` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `customer_id` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `place_id` varchar(255) NOT NULL,
  `detail` text NOT NULL,
  `lat` varchar(255) DEFAULT NULL,
  `long` varchar(255) DEFAULT NULL,
  `is_active` int(11) NOT NULL DEFAULT 0,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `admins`
--

CREATE TABLE `admins` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `surname` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `admins`
--

INSERT INTO `admins` (`id`, `name`, `surname`, `email`, `phone`, `password`, `created_at`, `updated_at`) VALUES
(1, 'Restaurant', 'Restaurant', 'no-reply@machhub.dk', '+4100000000', '$2y$10$G6e5vTIp8ShjbKN8Jsenx.i.52l3Z9Mq0M7RcAfyE/w.3QCCVXCvO', '2022-10-19 13:28:39', '2023-08-28 17:01:02');

-- --------------------------------------------------------

--
-- Table structure for table `attributes`
--

CREATE TABLE `attributes` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `value` longtext DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `attributes`
--

INSERT INTO `attributes` (`id`, `name`, `value`, `created_at`, `updated_at`) VALUES
(1, 'api_user_id', '0', '2023-08-28 17:01:01', '2023-08-28 17:01:01'),
(2, 'name', 'Restaurant', '2023-08-28 17:01:01', '2023-08-28 17:01:01'),
(3, 'email', 'no-reply@machhub.dk', '2023-08-28 17:01:01', '2023-08-28 17:01:01'),
(4, 'pickup_commission', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(5, 'delivery_commission', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(6, 'currier_commission', '0', '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(7, 'currier_type', 'none', '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(8, 'venue_id', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(9, 'currier_1', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(10, 'currier_2', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(11, 'currier_3', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(12, 'currier_4', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(13, 'currier_5', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(14, 'currier_6', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(15, 'currier_7', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(16, 'currier_8', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(17, 'currier_9', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(18, 'currier_10', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(19, 'currier_11', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(20, 'currier_12', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(21, 'currier_13', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(22, 'currier_14', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(23, 'terminal_type', 'none', '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(24, 'terminal_id', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(25, 'terminal_username', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02'),
(26, 'terminal_password', NULL, '2023-08-28 17:01:02', '2023-08-28 17:01:02');

-- --------------------------------------------------------

--
-- Table structure for table `campaigns`
--

CREATE TABLE `campaigns` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `title` varchar(255) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `image` varchar(255) NOT NULL,
  `detail` longtext DEFAULT NULL,
  `end_date` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `carts`
--

CREATE TABLE `carts` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `cart_id` varchar(255) NOT NULL,
  `product_id` int(11) NOT NULL,
  `qty` int(11) NOT NULL,
  `variants` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`variants`)),
  `extras` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`extras`)),
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `category_id` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `icon` varchar(255) NOT NULL,
  `banner` varchar(255) DEFAULT NULL,
  `sort` int(11) NOT NULL,
  `is_active` int(11) NOT NULL DEFAULT 0,
  `description` longtext DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `Sunday` varchar(255) DEFAULT NULL,
  `Monday` varchar(255) DEFAULT NULL,
  `Tuesday` varchar(255) DEFAULT NULL,
  `Wednesday` varchar(255) DEFAULT NULL,
  `Thursday` varchar(255) DEFAULT NULL,
  `Friday` varchar(255) DEFAULT NULL,
  `Saturday` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `coupons`
--

CREATE TABLE `coupons` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `code` varchar(255) NOT NULL,
  `status` int(11) NOT NULL,
  `start_date` timestamp NULL DEFAULT NULL,
  `end_date` timestamp NULL DEFAULT NULL,
  `max_use` int(11) NOT NULL,
  `user_max_use` int(11) NOT NULL,
  `total_use` int(11) NOT NULL DEFAULT 0,
  `sale_type` enum('1','2') NOT NULL,
  `sale_amount` double(8,2) NOT NULL,
  `min_cart_total` double(8,2) NOT NULL,
  `user_group` enum('1','2','3') DEFAULT '1',
  `user_id` int(11) DEFAULT NULL,
  `register_date` enum('1','2','3','4') DEFAULT '1',
  `order_source` enum('1','2','3') DEFAULT '1',
  `free_cargo` enum('0','1') DEFAULT '0',
  `order_count` int(11) NOT NULL DEFAULT 0,
  `delivery_type` int(11) NOT NULL DEFAULT 0,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `surname` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `gender` int(11) NOT NULL DEFAULT 0,
  `profile` varchar(255) DEFAULT 'default/customer.png',
  `is_verified` int(11) NOT NULL DEFAULT 0,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `remember_token` varchar(100) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `lang` varchar(255) NOT NULL DEFAULT 'da'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `customer_points`
--

CREATE TABLE `customer_points` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `customer_id` int(11) NOT NULL,
  `point` double NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `extras`
--

CREATE TABLE `extras` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `group_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `price` double NOT NULL,
  `is_disabled` tinyint(1) NOT NULL DEFAULT 0,
  `is_default` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `extra_groups`
--

CREATE TABLE `extra_groups` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `display_name` varchar(255) NOT NULL,
  `limit` varchar(255) NOT NULL,
  `order` int(11) NOT NULL DEFAULT 0,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `failed_jobs`
--

CREATE TABLE `failed_jobs` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `uuid` varchar(255) NOT NULL,
  `connection` text NOT NULL,
  `queue` text NOT NULL,
  `payload` longtext NOT NULL,
  `exception` longtext NOT NULL,
  `failed_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `jobs`
--

CREATE TABLE `jobs` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `queue` varchar(255) NOT NULL,
  `payload` longtext NOT NULL,
  `attempts` tinyint(3) UNSIGNED NOT NULL,
  `reserved_at` int(10) UNSIGNED DEFAULT NULL,
  `available_at` int(10) UNSIGNED NOT NULL,
  `created_at` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `migrations`
--

CREATE TABLE `migrations` (
  `id` int(10) UNSIGNED NOT NULL,
  `migration` varchar(255) NOT NULL,
  `batch` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `migrations`
--

INSERT INTO `migrations` (`id`, `migration`, `batch`) VALUES
(1, '2014_10_12_100000_create_password_resets_table', 1),
(2, '2019_08_19_000000_create_failed_jobs_table', 1),
(3, '2019_12_14_000001_create_personal_access_tokens_table', 1),
(4, '2022_10_16_165428_create_admins_table', 1),
(5, '2022_10_17_232442_create_settings_table', 1),
(9, '2022_10_19_163038_create_categories_table', 2),
(13, '2022_10_19_165802_create_products_table', 3),
(15, '2022_10_30_180611_create_phone_verifications_table', 5),
(16, '2022_11_05_232243_create_product_variants_table', 6),
(17, '2022_11_05_232344_create_product_variant_options_table', 6),
(25, '2022_11_10_213353_create_coupons_table', 11),
(26, '2022_11_12_165502_create_customer_points_table', 12),
(27, '2022_11_12_175013_create_point_actions_table', 13),
(28, '2022_11_12_224243_create_order_actions_table', 14),
(31, '2023_01_06_220032_create_tables_table', 16),
(33, '2023_01_06_220153_create_reservations_table', 17),
(35, '2023_01_12_145312_create_campaigns_table', 18),
(36, '2023_01_29_205058_create_extra_groups_table', 19),
(37, '2023_01_29_205356_create_extras_table', 19),
(38, '2023_01_29_220207_create_product_extra_groups_table', 20),
(39, '2022_11_06_155032_create_carts_table', 21),
(40, '2022_11_07_003212_create_order_products_table', 22),
(41, '2022_10_29_210321_create_customers_table', 23),
(43, '2022_12_08_010349_create_addresses_table', 24),
(44, '2023_03_08_193103_create_attributes_table', 25),
(45, '2023_03_08_194758_create_jobs_table', 26),
(46, '2022_11_07_002930_create_orders_table', 27),
(47, '2023_03_26_163519_add_products_type', 28),
(48, '2023_03_27_163603_update_product_type_migration', 28),
(49, '2023_04_01_003054_add_is_pre_order_column_orders', 28),
(50, '2023_04_01_183019_add_lang_customers_table', 28),
(51, '2023_04_23_000715_add_description_categories_table', 28),
(52, '2023_05_01_235916_add_default_and_disabled_field_extras_table', 29),
(53, '2023_05_08_121850_add_delivery_type_coupons_table', 29),
(54, '2023_05_29_203432_change_date_column_orders_table', 29),
(55, '2023_07_04_202800_category_extra', 29),
(56, '2023_07_07_183920_add_payment_id_column_orders_table', 29),
(57, '2023_08_04_185946_create_permission_tables', 29);

-- --------------------------------------------------------

--
-- Table structure for table `model_has_permissions`
--

CREATE TABLE `model_has_permissions` (
  `permission_id` bigint(20) UNSIGNED NOT NULL,
  `model_type` varchar(255) NOT NULL,
  `model_id` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `model_has_permissions`
--

INSERT INTO `model_has_permissions` (`permission_id`, `model_type`, `model_id`) VALUES
(1, 'App\\Models\\Admin', 1),
(2, 'App\\Models\\Admin', 1),
(3, 'App\\Models\\Admin', 1),
(4, 'App\\Models\\Admin', 1),
(5, 'App\\Models\\Admin', 1),
(6, 'App\\Models\\Admin', 1),
(7, 'App\\Models\\Admin', 1),
(8, 'App\\Models\\Admin', 1),
(9, 'App\\Models\\Admin', 1),
(10, 'App\\Models\\Admin', 1),
(11, 'App\\Models\\Admin', 1),
(12, 'App\\Models\\Admin', 1),
(13, 'App\\Models\\Admin', 1),
(14, 'App\\Models\\Admin', 1),
(15, 'App\\Models\\Admin', 1),
(16, 'App\\Models\\Admin', 1),
(17, 'App\\Models\\Admin', 1),
(18, 'App\\Models\\Admin', 1),
(19, 'App\\Models\\Admin', 1),
(20, 'App\\Models\\Admin', 1),
(21, 'App\\Models\\Admin', 1),
(22, 'App\\Models\\Admin', 1),
(23, 'App\\Models\\Admin', 1),
(24, 'App\\Models\\Admin', 1),
(25, 'App\\Models\\Admin', 1),
(26, 'App\\Models\\Admin', 1),
(27, 'App\\Models\\Admin', 1),
(28, 'App\\Models\\Admin', 1),
(29, 'App\\Models\\Admin', 1),
(30, 'App\\Models\\Admin', 1),
(31, 'App\\Models\\Admin', 1),
(32, 'App\\Models\\Admin', 1),
(33, 'App\\Models\\Admin', 1),
(34, 'App\\Models\\Admin', 1),
(35, 'App\\Models\\Admin', 1),
(36, 'App\\Models\\Admin', 1),
(37, 'App\\Models\\Admin', 1),
(38, 'App\\Models\\Admin', 1),
(39, 'App\\Models\\Admin', 1),
(40, 'App\\Models\\Admin', 1),
(41, 'App\\Models\\Admin', 1),
(42, 'App\\Models\\Admin', 1),
(43, 'App\\Models\\Admin', 1),
(44, 'App\\Models\\Admin', 1),
(45, 'App\\Models\\Admin', 1),
(46, 'App\\Models\\Admin', 1),
(47, 'App\\Models\\Admin', 1),
(48, 'App\\Models\\Admin', 1);

-- --------------------------------------------------------

--
-- Table structure for table `model_has_roles`
--

CREATE TABLE `model_has_roles` (
  `role_id` bigint(20) UNSIGNED NOT NULL,
  `model_type` varchar(255) NOT NULL,
  `model_id` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `order_number` varchar(255) NOT NULL,
  `payment_id` varchar(255) DEFAULT NULL,
  `customer_id` int(11) DEFAULT NULL,
  `customer` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`customer`)),
  `address` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`address`)),
  `prices` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`prices`)),
  `payment` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`payment`)),
  `status` int(11) DEFAULT NULL,
  `type` int(11) DEFAULT NULL,
  `note` longtext DEFAULT NULL,
  `other` longtext DEFAULT NULL,
  `date` datetime DEFAULT current_timestamp(),
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `is_pre_order` tinyint(1) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `order_actions`
--

CREATE TABLE `order_actions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `order_id` int(11) NOT NULL,
  `text` text NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `order_products`
--

CREATE TABLE `order_products` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `order_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `qty` int(11) NOT NULL,
  `unit_price` double NOT NULL,
  `price` double NOT NULL,
  `variants` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`variants`)),
  `extras` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`extras`)),
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `password_resets`
--

CREATE TABLE `password_resets` (
  `email` varchar(255) NOT NULL,
  `token` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `permissions`
--

CREATE TABLE `permissions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `guard_name` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `permissions`
--

INSERT INTO `permissions` (`id`, `name`, `guard_name`, `created_at`, `updated_at`) VALUES
(1, 'show product', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(2, 'create product', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(3, 'edit product', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(4, 'delete product', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(5, 'show category', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(6, 'create category', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(7, 'edit category', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(8, 'delete category', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(9, 'show extra', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(10, 'create extra', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(11, 'edit extra', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(12, 'delete extra', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(13, 'show order', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(14, 'approve order', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(15, 'cancel order', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(16, 'call currier order', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(17, 'ship order', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(18, 'complete order', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(19, 'edit order', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(20, 'show coupon', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(21, 'create coupon', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(22, 'edit coupon', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(23, 'delete coupon', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(24, 'show campaign', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(25, 'create campaign', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(26, 'edit campaign', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(27, 'delete campaign', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(28, 'show user', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(29, 'create user', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(30, 'edit user', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(31, 'delete user', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(32, 'show desk', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(33, 'create desk', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(34, 'delete desk', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(35, 'show reservation', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(36, 'delete reservation', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(37, 'show reports', 'admin', '2023-08-28 17:01:04', '2023-08-28 17:01:04'),
(38, 'edit site settings', 'admin', '2023-08-28 17:01:05', '2023-08-28 17:01:05'),
(39, 'edit order settings', 'admin', '2023-08-28 17:01:05', '2023-08-28 17:01:05'),
(40, 'edit hour settings', 'admin', '2023-08-28 17:01:05', '2023-08-28 17:01:05'),
(41, 'edit qr settings', 'admin', '2023-08-28 17:01:05', '2023-08-28 17:01:05'),
(42, 'edit theme settings', 'admin', '2023-08-28 17:01:05', '2023-08-28 17:01:05'),
(43, 'edit language settings', 'admin', '2023-08-28 17:01:05', '2023-08-28 17:01:05'),
(44, 'show market', 'admin', '2023-08-28 17:01:05', '2023-08-28 17:01:05'),
(45, 'show admin', 'admin', '2023-08-28 17:01:05', '2023-08-28 17:01:05'),
(46, 'create admin', 'admin', '2023-08-28 17:01:05', '2023-08-28 17:01:05'),
(47, 'edit admin', 'admin', '2023-08-28 17:01:05', '2023-08-28 17:01:05'),
(48, 'delete admin', 'admin', '2023-08-28 17:01:05', '2023-08-28 17:01:05');

-- --------------------------------------------------------

--
-- Table structure for table `personal_access_tokens`
--

CREATE TABLE `personal_access_tokens` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `tokenable_type` varchar(255) NOT NULL,
  `tokenable_id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `token` varchar(64) NOT NULL,
  `abilities` text DEFAULT NULL,
  `last_used_at` timestamp NULL DEFAULT NULL,
  `expires_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `phone_verifications`
--

CREATE TABLE `phone_verifications` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `customer_id` int(11) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `code` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `point_actions`
--

CREATE TABLE `point_actions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `customer_id` int(11) NOT NULL,
  `order_id` int(11) NOT NULL,
  `amount` double NOT NULL,
  `type` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `category_id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `image` varchar(255) NOT NULL,
  `without_discount_price` double NOT NULL,
  `price` double NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `sort` int(11) NOT NULL DEFAULT 0,
  `is_active` int(11) NOT NULL,
  `keywords` longtext DEFAULT NULL,
  `materials` longtext DEFAULT NULL,
  `discount` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`discount`)),
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `product_extra_groups`
--

CREATE TABLE `product_extra_groups` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `product_id` int(11) NOT NULL,
  `product_variant_option_id` int(11) NOT NULL,
  `extra_group_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `product_variants`
--

CREATE TABLE `product_variants` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `product_id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `product_variant_options`
--

CREATE TABLE `product_variant_options` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `product_variant_id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `price` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `reservations`
--

CREATE TABLE `reservations` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `customer_id` int(11) NOT NULL,
  `table_id` int(11) NOT NULL,
  `person` int(11) NOT NULL DEFAULT 0,
  `date` varchar(255) NOT NULL,
  `time` varchar(255) NOT NULL,
  `status` int(11) NOT NULL DEFAULT 0,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `guard_name` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `role_has_permissions`
--

CREATE TABLE `role_has_permissions` (
  `permission_id` bigint(20) UNSIGNED NOT NULL,
  `role_id` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `settings`
--

CREATE TABLE `settings` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `value` longtext DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `settings`
--

INSERT INTO `settings` (`id`, `name`, `value`, `created_at`, `updated_at`) VALUES
(1, 'name', 'Waiters', '2022-10-20 17:23:47', '2022-10-20 17:23:47'),
(4, 'logo', 'settings/tDLTG5ceDHkHFVvCw9NzxmJgub41glNQQ5quWWFO.png', '2022-10-20 17:23:48', '2023-01-05 14:29:11'),
(5, 'title', 'Restaurant', '2022-10-20 17:28:40', '2023-04-25 12:48:35'),
(6, 'favicon', 'settings/iGqWnUro7UUTlzR1p20HupQIhHYQYj1kjjOJU6T3.png', '2022-10-20 17:28:40', '2022-10-20 17:28:40'),
(7, 'activeDays', 'Pazartesi,Salı,Çarşamba,Perşembe,Cuma,Cumartesi,Pazar', '2022-10-20 17:53:23', '2023-03-14 20:39:28'),
(8, 'order_status', '0', '2022-10-20 17:53:23', '2023-04-25 12:48:41'),
(9, 'min_order_amount', '0', '2022-10-20 17:53:23', '2023-08-28 17:06:44'),
(10, 'delivery_time', '30-35 m', '2022-10-20 17:53:23', '2022-10-30 16:16:23'),
(11, 'open_time', '09:00', '2022-10-20 18:07:49', '2022-10-31 12:06:26'),
(12, 'close_time', '01:25', '2022-10-20 18:07:49', '2023-02-09 22:24:46'),
(13, 'restaurant_location_lat', '37.1804373', '2022-10-20 18:26:56', '2023-02-01 15:36:28'),
(14, 'restaurant_location_long', '33.2353032', '2022-10-20 18:26:56', '2023-02-01 15:36:28'),
(15, 'description', NULL, '2022-10-30 13:35:50', '2022-10-30 13:35:50'),
(16, 'banner', 'settings/vMBDFuFbePl6L48lmCGxrfud3ac0I1dJ1yuCscUX.svg', '2022-10-30 13:35:50', '2022-10-30 13:35:50'),
(17, 'address', 'København', '2022-11-06 18:47:00', '2023-02-01 15:36:27'),
(18, 'delivery_price', '50', '2022-11-06 20:27:07', '2023-08-28 17:06:45'),
(19, 'free_delivery_limit', '0', '2022-11-06 20:27:07', '2023-08-28 17:06:45'),
(20, 'order_bonus', '0', '2022-11-12 13:59:43', '2023-08-28 17:06:45'),
(21, 'currier_type', 'MANUAL', '2023-01-05 12:28:12', '2023-08-28 17:06:44'),
(24, 'template_module_search_title', 'Machhub', '2023-01-12 10:34:35', '2023-02-01 13:06:17'),
(25, 'template_module_search_detail', 'Lorem ipsum dolor, sit amet consectetur adipisicing elit. Dolorem blanditiis, sequi at beatae fugiat nobis quasi nihil ut vitae! Dolores.', '2023-01-12 10:34:35', '2023-01-12 10:34:35'),
(26, 'template_module_search_input_text', 'Aramaya Başla', '2023-01-12 10:36:03', '2023-01-12 10:36:03'),
(27, 'template_module_search_image', 'settings/H4P14xfM4ljmBv3JuCEr0WTBElp6GAo3wOufmEao.png', '2023-01-12 10:36:03', '2023-01-30 21:33:55'),
(28, 'template_module_search_is_active', '1', '2023-01-12 10:37:34', '2023-01-19 11:21:48'),
(29, 'template_module_slider_is_active', '0', '2023-01-12 10:47:50', '2023-03-05 12:31:47'),
(30, 'template_module_slider_image', 'settings/3GWCdg1VqlFQVKY0Q2pPJ9yOwKT61te12pOFdGX1.png', '2023-01-12 10:47:50', '2023-02-18 12:34:36'),
(31, 'template_module_slider_button_text', 'Cart', '2023-01-12 10:50:27', '2023-02-18 12:34:30'),
(32, 'template_module_slider_title', 'Slider', '2023-01-12 10:53:40', '2023-01-12 10:53:40'),
(33, 'template_module_slider_button_url', 'cart.index', '2023-01-12 10:53:40', '2023-01-12 11:18:04'),
(34, 'template_module_category_length', '999', '2023-01-12 11:01:38', '2023-01-19 11:26:00'),
(35, 'template_module_category_active_color', '#06cb52', '2023-01-12 11:01:38', '2023-01-12 11:08:54'),
(36, 'template_module_category_is_active', '1', '2023-01-12 11:07:16', '2023-01-12 11:07:16'),
(37, 'template_module_products_is_active', '1', '2023-01-12 11:07:16', '2023-01-12 11:10:17'),
(38, 'template_module_product_length', '999', '2023-01-12 11:09:19', '2023-01-12 11:10:23'),
(39, 'template_module_products_show_sale_price', '1', '2023-01-12 11:12:02', '2023-01-12 11:12:02'),
(40, 'template_module_products_show_time', '1', '2023-01-12 11:12:10', '2023-01-12 11:12:10'),
(41, 'template_module_slider_show_no_discount_price', '1', '2023-01-12 11:12:20', '2023-01-12 11:12:20'),
(42, 'template_module_products_show_comment_count', '0', '2023-01-12 11:13:50', '2023-01-12 11:13:50'),
(43, 'template_module_products_show_no_discount_price', '1', '2023-01-12 11:13:50', '2023-01-19 11:35:18'),
(44, 'template_module_products_show_type', 'popup', '2023-01-12 11:18:04', '2023-01-29 20:36:56'),
(45, 'template_module_campaign_is_active', '1', '2023-01-12 11:51:25', '2023-01-12 12:19:12'),
(46, 'template_module_featured_products_is_active', '1', '2023-01-12 11:51:25', '2023-01-12 11:51:31'),
(47, 'template_module_featured_products_length', '6', '2023-01-12 11:51:25', '2023-01-12 12:25:46'),
(48, 'template_module_featured_products_title', 'Featured ', '2023-01-12 11:51:25', '2023-01-12 11:51:25'),
(49, 'template_module_campaign_title', 'Campains', '2023-01-12 12:24:45', '2023-01-12 12:24:45'),
(50, 'phone', NULL, '2023-01-29 16:34:38', '2023-04-25 12:48:35'),
(51, 'template_module_general_background', '#151a21', '2023-02-01 12:48:40', '2023-02-01 13:06:03'),
(52, 'template_module_general_header_background', '#151a21', '2023-02-01 12:52:57', '2023-02-01 13:06:03'),
(53, 'template_module_products_box_background', '#ffffff', '2023-02-01 12:54:38', '2023-02-01 12:58:14'),
(54, 'template_module_search_input_background', '#56279a20', '2023-02-01 13:00:30', '2023-02-01 13:06:57'),
(55, 'template_module_general_background', '#151a21', '2023-02-01 13:17:17', '2023-02-01 13:17:17'),
(56, 'template_module_general_header_background', '#151a21', '2023-02-01 13:17:17', '2023-02-01 13:17:17'),
(57, 'template_module_products_box_background', '#ffffff', '2023-02-01 13:17:17', '2023-02-01 13:17:17'),
(58, 'template_module_search_input_background', '#56279a20', '2023-02-01 13:17:17', '2023-02-01 13:17:17'),
(59, 'template_module_general_background', '#151a21', '2023-02-01 15:20:21', '2023-02-01 15:20:21'),
(60, 'template_module_general_header_background', '#151a21', '2023-02-01 15:20:21', '2023-02-01 15:20:21'),
(61, 'template_module_products_box_background', '#ffffff', '2023-02-01 15:20:21', '2023-02-01 15:20:21'),
(62, 'template_module_search_input_background', '#56279a20', '2023-02-01 15:20:21', '2023-02-01 15:20:21'),
(65, 'delivery_area', '10', '2023-02-01 15:48:21', '2023-08-28 17:06:45'),
(66, 'paid_delivery_start', '5', '2023-02-01 16:42:55', '2023-02-01 16:42:55'),
(67, 'km_delivery_price', '10', '2023-02-01 16:42:55', '2023-08-28 17:06:45'),
(68, 'default_product_image', 'settings/0kYNsCArVscygWhwWZ5GNHMVIVny7govNFKTb8y8.png', '2023-02-01 18:47:04', '2023-04-25 12:48:22'),
(69, 'time_Monday_open', '08:00-03:00', '2023-02-10 21:33:05', '2023-02-10 23:25:31'),
(70, 'time_Monday_delivery', '10:00-03:00', '2023-02-10 21:33:05', '2023-03-05 20:55:54'),
(71, 'time_Monday_pickup', '09:00-03:00', '2023-02-10 21:33:05', '2023-02-10 23:25:31'),
(72, 'time_Tuesday_open', '08:00-23:00', '2023-02-10 21:33:05', '2023-03-14 20:40:18'),
(73, 'time_Tuesday_delivery', '10:00-22:00', '2023-02-10 21:33:05', '2023-02-10 23:25:31'),
(74, 'time_Tuesday_pickup', '09:00-03:00', '2023-02-10 21:33:05', '2023-02-10 23:25:31'),
(75, 'time_Wednesday_open', '08:00-03:00', '2023-02-10 21:33:05', '2023-02-10 23:25:31'),
(76, 'time_Wednesday_delivery', '10:00-22:00', '2023-02-10 21:33:05', '2023-02-10 23:25:31'),
(77, 'time_Wednesday_pickup', '09:00-03:00', '2023-02-10 21:33:05', '2023-02-10 23:25:31'),
(78, 'time_Thursday_open', '08:00-03:00', '2023-02-10 21:33:05', '2023-02-10 23:25:31'),
(79, 'time_Thursday_delivery', '10:00-22:00', '2023-02-10 21:33:05', '2023-02-10 23:25:31'),
(80, 'time_Thursday_pickup', '09:00-03:00', '2023-02-10 21:33:05', '2023-02-10 23:25:31'),
(81, 'time_Friday_open', '08:00-03:00', '2023-02-10 21:33:05', '2023-02-10 23:25:31'),
(82, 'time_Friday_delivery', '10:00-22:00', '2023-02-10 21:33:05', '2023-02-10 23:25:31'),
(83, 'time_Friday_pickup', '09:00-03:00', '2023-02-10 21:33:05', '2023-02-10 23:25:31'),
(84, 'time_Saturday_open', '08:00-03:00', '2023-02-10 21:33:05', '2023-02-10 23:25:31'),
(85, 'time_Saturday_delivery', '10:00-22:00', '2023-02-10 21:33:05', '2023-02-11 12:33:05'),
(86, 'time_Saturday_pickup', '09:00-22:00', '2023-02-10 21:33:05', '2023-02-11 00:15:40'),
(87, 'time_Sunday_open', '09:00-18:00', '2023-02-10 21:33:05', '2023-02-11 10:56:53'),
(88, 'time_Sunday_delivery', '10:00-03:00', '2023-02-10 21:33:05', '2023-03-05 20:55:54'),
(89, 'time_Sunday_pickup', '09:00-22:00', '2023-02-10 21:33:05', '2023-02-11 00:15:40'),
(90, 'max_order_day', '7', '2023-02-11 00:04:22', '2023-02-11 00:04:22'),
(91, 'template_module_search_input_text_color', '#e20808', '2023-02-11 15:40:17', '2023-02-11 15:40:17'),
(92, 'template_module_footer_campaign_is_active', '1', '2023-02-18 12:32:10', '2023-02-18 12:38:34'),
(93, 'template_module_footer_campaign_title', 'Machhub', '2023-02-18 12:32:10', '2023-02-22 21:53:52'),
(94, 'template_module_footer_campaign_button_url', 'customer.reservation.create', '2023-02-18 12:32:10', '2023-02-18 12:32:10'),
(95, 'template_module_footer_campaign_button_text', 'Reservation', '2023-02-18 12:34:01', '2023-02-18 12:34:01'),
(96, 'template_module_footer_campaign_image', 'settings/vIYjugrtWHgMs3GhZHGF3tvyyPAa4XVmtfukhdMi.jpg', '2023-02-18 12:34:01', '2023-02-18 12:34:20'),
(97, 'template_module_footer_campaign_detail', 'Quibus occurrere bene pertinax miles explicatis ordinibus parans hastisque feriens scuta qui habitus iram. Bene pertinax miles explicatis ordinibus.', '2023-02-18 12:35:41', '2023-02-18 12:35:41'),
(98, 'order_delivery_time', '60', '2023-03-09 15:23:52', '2023-08-28 17:06:44'),
(99, 'order_pickup_time', '15', '2023-03-09 15:23:52', '2023-08-28 17:06:44'),
(100, 'order_restaurant_time', '45', '2023-03-09 15:23:52', '2023-08-28 17:06:44'),
(101, 'theme', 'traditional', '2023-03-13 15:26:01', '2023-03-22 09:22:02'),
(102, 'traditional_background_image', 'settings/UknDObExbX0tVtojhBDv4JUmrtPuqi5QJZurnlTu.png', '2023-04-25 12:48:22', '2023-04-25 12:48:22'),
(103, 'email', NULL, '2023-04-25 12:48:35', '2023-04-25 12:48:35'),
(104, 'health_report_url', NULL, '2023-04-25 12:48:35', '2023-04-25 12:48:35'),
(105, 'facebook', NULL, '2023-04-25 12:48:35', '2023-04-25 12:48:35'),
(106, 'twitter', NULL, '2023-04-25 12:48:35', '2023-04-25 12:48:35'),
(107, 'instagram', NULL, '2023-04-25 12:48:36', '2023-04-25 12:48:36'),
(108, 'header_tags', NULL, '2023-04-25 12:48:36', '2023-04-25 12:48:36'),
(109, 'footer_tags', NULL, '2023-04-25 12:48:36', '2023-04-25 12:48:36'),
(110, 'active_pickup', '1', '2023-08-28 17:06:44', '2023-08-28 17:06:44'),
(111, 'active_delivery', '1', '2023-08-28 17:06:44', '2023-08-28 17:06:44'),
(112, 'active_table', '1', '2023-08-28 17:06:44', '2023-08-28 17:06:44'),
(113, 'min_order_amount_takeaway', '0', '2023-08-28 17:06:44', '2023-08-28 17:06:44'),
(114, 'order_delivery_pickup_time', '20', '2023-08-28 17:06:44', '2023-08-28 17:06:44'),
(115, 'paymentfee_delivery', '5', '2023-08-28 17:06:44', '2023-08-28 17:06:44'),
(116, 'paymentfee_table', '3', '2023-08-28 17:06:44', '2023-08-28 17:06:44'),
(117, 'paymentfee_other', '3', '2023-08-28 17:06:44', '2023-08-28 17:06:44'),
(118, 'bagfee', '4', '2023-08-28 17:06:45', '2023-08-28 17:06:45'),
(119, 'city', NULL, '2023-08-28 17:06:45', '2023-08-28 17:06:45');

-- --------------------------------------------------------

--
-- Table structure for table `tables`
--

CREATE TABLE `tables` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `addresses`
--
ALTER TABLE `addresses`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `admins`
--
ALTER TABLE `admins`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `admins_email_unique` (`email`);

--
-- Indexes for table `attributes`
--
ALTER TABLE `attributes`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `campaigns`
--
ALTER TABLE `campaigns`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `carts`
--
ALTER TABLE `carts`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `coupons`
--
ALTER TABLE `coupons`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `customer_points`
--
ALTER TABLE `customer_points`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `extras`
--
ALTER TABLE `extras`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `extra_groups`
--
ALTER TABLE `extra_groups`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `failed_jobs`
--
ALTER TABLE `failed_jobs`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `failed_jobs_uuid_unique` (`uuid`);

--
-- Indexes for table `jobs`
--
ALTER TABLE `jobs`
  ADD PRIMARY KEY (`id`),
  ADD KEY `jobs_queue_index` (`queue`);

--
-- Indexes for table `migrations`
--
ALTER TABLE `migrations`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `model_has_permissions`
--
ALTER TABLE `model_has_permissions`
  ADD PRIMARY KEY (`permission_id`,`model_id`,`model_type`),
  ADD KEY `model_has_permissions_model_id_model_type_index` (`model_id`,`model_type`);

--
-- Indexes for table `model_has_roles`
--
ALTER TABLE `model_has_roles`
  ADD PRIMARY KEY (`role_id`,`model_id`,`model_type`),
  ADD KEY `model_has_roles_model_id_model_type_index` (`model_id`,`model_type`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `order_actions`
--
ALTER TABLE `order_actions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `order_products`
--
ALTER TABLE `order_products`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `password_resets`
--
ALTER TABLE `password_resets`
  ADD KEY `password_resets_email_index` (`email`);

--
-- Indexes for table `permissions`
--
ALTER TABLE `permissions`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `permissions_name_guard_name_unique` (`name`,`guard_name`);

--
-- Indexes for table `personal_access_tokens`
--
ALTER TABLE `personal_access_tokens`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `personal_access_tokens_token_unique` (`token`),
  ADD KEY `personal_access_tokens_tokenable_type_tokenable_id_index` (`tokenable_type`,`tokenable_id`);

--
-- Indexes for table `phone_verifications`
--
ALTER TABLE `phone_verifications`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `point_actions`
--
ALTER TABLE `point_actions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `product_extra_groups`
--
ALTER TABLE `product_extra_groups`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `product_variants`
--
ALTER TABLE `product_variants`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `product_variant_options`
--
ALTER TABLE `product_variant_options`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `reservations`
--
ALTER TABLE `reservations`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `roles_name_guard_name_unique` (`name`,`guard_name`);

--
-- Indexes for table `role_has_permissions`
--
ALTER TABLE `role_has_permissions`
  ADD PRIMARY KEY (`permission_id`,`role_id`),
  ADD KEY `role_has_permissions_role_id_foreign` (`role_id`);

--
-- Indexes for table `settings`
--
ALTER TABLE `settings`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tables`
--
ALTER TABLE `tables`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `addresses`
--
ALTER TABLE `addresses`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `admins`
--
ALTER TABLE `admins`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `attributes`
--
ALTER TABLE `attributes`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=27;

--
-- AUTO_INCREMENT for table `campaigns`
--
ALTER TABLE `campaigns`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `carts`
--
ALTER TABLE `carts`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `categories`
--
ALTER TABLE `categories`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `coupons`
--
ALTER TABLE `coupons`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `customer_points`
--
ALTER TABLE `customer_points`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `extras`
--
ALTER TABLE `extras`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `extra_groups`
--
ALTER TABLE `extra_groups`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `failed_jobs`
--
ALTER TABLE `failed_jobs`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `jobs`
--
ALTER TABLE `jobs`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `migrations`
--
ALTER TABLE `migrations`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=58;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `order_actions`
--
ALTER TABLE `order_actions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `order_products`
--
ALTER TABLE `order_products`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `permissions`
--
ALTER TABLE `permissions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=49;

--
-- AUTO_INCREMENT for table `personal_access_tokens`
--
ALTER TABLE `personal_access_tokens`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `phone_verifications`
--
ALTER TABLE `phone_verifications`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `point_actions`
--
ALTER TABLE `point_actions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `product_extra_groups`
--
ALTER TABLE `product_extra_groups`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `product_variants`
--
ALTER TABLE `product_variants`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `product_variant_options`
--
ALTER TABLE `product_variant_options`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `reservations`
--
ALTER TABLE `reservations`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `roles`
--
ALTER TABLE `roles`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `settings`
--
ALTER TABLE `settings`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=120;

--
-- AUTO_INCREMENT for table `tables`
--
ALTER TABLE `tables`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `model_has_permissions`
--
ALTER TABLE `model_has_permissions`
  ADD CONSTRAINT `model_has_permissions_permission_id_foreign` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `model_has_roles`
--
ALTER TABLE `model_has_roles`
  ADD CONSTRAINT `model_has_roles_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `role_has_permissions`
--
ALTER TABLE `role_has_permissions`
  ADD CONSTRAINT `role_has_permissions_permission_id_foreign` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `role_has_permissions_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
