export module ModuleA;

// Define keyword to macro function call feature
// Example : #define <macro name> <macro feature value>
#define TEST 42

export namespace ModuleA_1_NS {
	// Create feature code function
	int fn_internal() {
		return TEST;
	}

	// Export to created feature code function
	export int fn() {
		return fn_internal();
	}
}